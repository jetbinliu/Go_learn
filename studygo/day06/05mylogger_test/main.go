package main

import (
	"go_git/Go_learn/studygo/day06/mylogger"
)

// 测试自己写的日志库

func main() {
	log := mylogger.NewLog("Debug")
	log.Debug("log debug")
	//time.Sleep(time.Second * 2)
	log.Info("log info")
	log.Error("log error")
}
