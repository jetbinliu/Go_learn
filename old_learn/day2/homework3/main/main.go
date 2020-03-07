package main

import (
	"fmt"
)

func fact(n int) {
	var a int = 1
	var sum int = 0
	for i := 1; i <= n ; i++{
		a = a * i
		sum = sum + a
	}
	fmt.Println(sum)
}

func main() {
    fact(5)
}