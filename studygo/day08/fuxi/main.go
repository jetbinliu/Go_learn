package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel

// 往通道中发送值
func sendNum(ch chan<- int) {
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(time.Second * 1)
	}
}

func main() {

	ch := make(chan int, 2)
	// ch <- 100
	// <-ch
	// ch <- 200
	// //<-ch
	// x, ok := <-ch
	// fmt.Println(x, ok)
	go sendNum(ch)
	for {
		x, ok := <-ch
		fmt.Println(x, ok)
		time.Sleep(time.Second * 2)
	}
}
