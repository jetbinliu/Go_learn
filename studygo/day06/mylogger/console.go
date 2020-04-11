package mylogger

import (
	"fmt"
	"time"
)

// 往终端写入日志相关内容

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func Ptime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func (l Logger) enable(LogLevel LogLevel) bool {
	return LogLevel >= l.Level
}

func log(lv LogLevel, msg string) {
	funcName, FileName, lineNo := getInfo(3)
	fmt.Printf("[%s][%s][%s:%s:%d] %s\n", Ptime(), getLogString(lv), FileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}
