package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {

	for {
		Second := time.Now().Unix()
		fmt.Println(Second)
		if Second%Female == 0 {
			fmt.Println(Female)
		} else {
			fmt.Println(Man)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
