package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	// lock.Lock()
	// x++
	atomic.AddInt64(&x, 1)
	wg.Done()

	// lock.Unlock()
}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
