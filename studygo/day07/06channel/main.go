package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	//fmt.Println(b)     // nil
	b = make(chan int) // 通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("x:", x)
	}()
	b <- 10
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
	close(b)
}

func bufChannel() {
	//fmt.Println(b)         // nil
	c := make(chan int, 1) // 通道的初始化
	c <- 1000
	x := <-c

	fmt.Println("10发送到通道c中了...x:", x)
	close(c)
}

func main() {
	bufChannel()
	noBufChannel()
}
