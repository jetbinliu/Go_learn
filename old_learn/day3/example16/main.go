package main

import (
	"fmt"
)

func test (a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}//(a,b)
	//return result
	return result(a,b)
}

func main() {
	fmt.Println(test(100,200))
	var i int = 0
	defer fmt.Println(i)

	i = 10
	defer fmt.Println(i)

	i = 11
	defer fmt.Println(i)

	i = 12
	fmt.Println(i)
}