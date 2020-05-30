package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么要context?

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
	for {
		fmt.Println("张三")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break
		default:

		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	exitChan <- true
	wg.Wait()
}
