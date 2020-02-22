package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"sync/atomic"
)

var lock sync.Mutex
var rwLock sync.RWMutex


func testMap() {
	a := make(map[int]int,10)
	a[1] = 11
	a[2] = 12
	a[3] = 13

	for i := 0; i < 2; i++{
		go func(b map[int]int){
			lock.Lock()
			b[1] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}

func testRWLock() {
	var count int32
	a := make(map[int]int,10)
	a[1] = 11
	a[2] = 12
	a[3] = 13

	for i := 0; i < 2; i++{
		go func(b map[int]int){
			//rwLock.Lock()
			lock.Lock()
			b[1] = rand.Intn(100)
			time.Sleep(10*time.Millisecond)
			//rwLock.Unlock()
			lock.Unlock()
		}(a)
	}

	for m := 0; m < 100; m++{
		go func(b map[int]int){
			for {
			//rwLock.RLock()
			lock.Lock()
			time.Sleep(time.Millisecond)
			//fmt.Println(a)
			//rwLock.RUnlock()
			lock.Unlock()
			atomic.AddInt32(&count,1)
		}
		}(a)
	}
	time.Sleep(2*time.Second)
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	testRWLock()
}