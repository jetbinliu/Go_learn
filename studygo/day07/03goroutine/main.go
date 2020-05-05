package main

import (
	"fmt"
	"time"
)

// goroutine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		//go hello(i) // 开启一个单独的goroutine去执行hello函数（任务）
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("main")
	// main函数结束了，由main函数启动的goroutine也都结束了
}
