package main

import (
	"fmt"
	"go_git/Go_learn/day1/goroute_example/goroute"
)

func main(){

	c := make(chan int,1)
	goroute.Add(1,2,c)
	s := <- c
    fmt.Println(s)
}