package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 往文件里面写日志相关代码
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
	logChan     chan *logMsg
}

type logMsg struct {
	Level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewFileLogger 构造函数
func NewFileLogger(LevelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(LevelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, 50000),
	}
	err = fl.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile() (err error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v \n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err file failed,err:%v \n", err)
		return err
	}
	// 日志文件都打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj

	// 开启1个后台的goroutine去往文件里写日志，开多个有问题
	go f.writeLogBackground()
	return nil

}

// 判断是否需要记录该日志
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

// 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	// 如果当前文件大小 >= 日志文件的最大值，则返回true
	return fileInfo.Size() >= f.maxFileSize

}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件
	nowStr := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n ", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) // 拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	// 1.关闭当前日志文件
	file.Close()
	// 2.备份一下 rename  xxx.log --> xx.log.bak201908031709
	os.Rename(logName, newLogName)
	// 3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file faild,err :%v\n", err)
		return nil, err
	}
	// 4.将打开的新日志文件对象赋值个f.fileObj
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) // 日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			// 把日志拼出来
			logInfo := fmt.Sprintf("[%s][%s][%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.Level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.Level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj) // 日志文件
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				// 如果要记录的日志大于等于ERROR级别，还要在ERROR log日志文件中再记录一份
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志先休息500ms
			time.Sleep(time.Millisecond * 500)
		}
	}

}

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, FileName, lineNo := getInfo(3)
		// 先把日志发送到通道中
		logTmp := &logMsg{
			Level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  FileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default: // 把日志丢掉，保证程序不出现阻塞
		}

	}
}

// Debug ...
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info ...
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning ...
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error ...
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Close ...
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
