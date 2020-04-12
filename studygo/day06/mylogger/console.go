package mylogger

import (
	"fmt"
	"time"
)

// 往终端写入日志相关内容

// ConsloeLogger 日志结构体
type ConsloeLogger struct {
	Level LogLevel
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) ConsloeLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsloeLogger{
		Level: level,
	}
}

func (l ConsloeLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= l.Level
}

// Ptime ...
func Ptime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func (l ConsloeLogger) log(lv LogLevel, format string, a ...interface{}) {
	if l.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		funcName, FileName, lineNo := getInfo(3)
		fmt.Printf("[%s][%s][%s:%s:%d] %s\n", Ptime(), getLogString(lv), FileName, funcName, lineNo, msg)
	}
}

// Debug ...
func (l ConsloeLogger) Debug(format string, a ...interface{}) {
	l.log(DEBUG, format, a...)
}

// Info ...
func (l ConsloeLogger) Info(format string, a ...interface{}) {
	l.log(INFO, format, a...)
}

// Warning ...
func (l ConsloeLogger) Warning(format string, a ...interface{}) {
	l.log(WARNING, format, a...)
}

// Debug ...
func (l ConsloeLogger) Error(format string, a ...interface{}) {
	l.log(ERROR, format, a...)
}
