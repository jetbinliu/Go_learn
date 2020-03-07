package main

import (
	"fmt"
	"time"
)

func test() {
	// s1 := new([]int)
	// fmt.Println(s1)

	// s2 := make([]int,10)
	// fmt.Println(s2)

	// *s1 = make([]int,5)
	// (*s1)[0] = 100
	// s2[0] = 100
	// fmt.Println(s1)
	// return
	s1 := new(int)
	*s1 = 6
	fmt.Println(*s1)
	//s2 := make()

}

func main() {
	fmt.Println("hello")
	time.Sleep(time.Second)
	main()

}