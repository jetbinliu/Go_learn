package main

import (
	"fmt"
	"time"
)

func main() {

	startTime := time.Now().UnixNano()

	fmt.Println("stattime:",startTime)
	time.Sleep(time.Second)
	now := time.Now()
	fmt.Println(now.Format("2006/1/02 15:04:05"))

	endTime := time.Now().UnixNano()
	fmt.Println("endtime:",endTime)

	diffTimeMicro := (endTime - startTime)/1000/1000
	fmt.Println("diffTime:",diffTimeMicro)
}