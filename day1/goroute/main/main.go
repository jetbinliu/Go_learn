package main

import (
	"go_git/Go_learn/day1/goroute/route"
	"time"
)

func main() {

	//var chain chan int
	chain := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		go time.Sleep(40*time.Second)
		chain <- i
		route.Route(chain)
		
	}
}