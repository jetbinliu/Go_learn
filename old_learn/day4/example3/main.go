package main

import (
	"fmt"
	"time"
)

func recursive(n int) {
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n > 10 {
		return
	}
	fmt.Println(n)
	recursive(n+1)
}

func factor(n int) int {
	fmt.Println(n)
	if (n == 1){
		return 1
	}
	
	return factor(n-1) * n
}

func fibonacci(n int) int{
	if n <= 1{
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
	//recursive(0)
	//result := factor(5)
	//fmt.Println(result)
	// for i := 0; i < 10; i++{
	// 	fmt.Println(fibonacci(i))
	// }
	fmt.Println(fibonacci(3))
}