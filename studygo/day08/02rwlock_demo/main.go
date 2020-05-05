package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var (
	x    = 0
	wg   sync.WaitGroup
	lock sync.Mutex
	// lock sync.RWMutex
)

func read() {
	defer wg.Done()
	lock.Lock()
	// lock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.RUnlock()
	lock.Unlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	wg.Wait()
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
