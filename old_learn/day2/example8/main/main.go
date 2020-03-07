package main

import (
	"fmt"
)

func modify( a int ){
	a = 20
}

func modify1( c *int ) {
	*c = 20
}

func main() {

	var a int = 10
	var b chan int = make(chan int,1)
	fmt.Println("a=",a)
	fmt.Println("b=",b)

	modify(a)
	fmt.Println("a=",a)

	modify1(&a)
	fmt.Println("a=",a)
}