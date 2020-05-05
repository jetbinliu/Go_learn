package main

import (
	"go_git/Go_learn/studygo/day08/mylogger"
)

//var log mylogger.Logger // 声明一个接口变量

// 测试自己写的日志库

func main() {
	// 输出到终端
	//log := mylogger.NewConsoleLogger("Info")

	// 输入到文件
	log := mylogger.NewFileLogger("INFO", "./", "test.log", 10*1024*1024) // 10MB

	id := 1000
	name := "zhangsan"
	for {
		log.Debug("log debug,id:%d ,name:%s", id, name) // 类似于fmt.Println
		//time.Sleep(time.Second * 2)
		log.Info("log info id:%d,name:%s", id, name)
		log.Error("log error id:%d ,name:%s", id, name)
		// time.Sleep(time.Second)
	}
}
