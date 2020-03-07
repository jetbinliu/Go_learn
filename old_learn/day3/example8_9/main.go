package main

import (
	"fmt"
)

func pointer(i *int) {
	*i = 3
	fmt.Println(*i)
}

func main() {
	a := 123
	fmt.Println("a is:", a, "a's address is :", &a)
	var p *int
	//a = 12345
	p = &a
	//*p = 12345
	fmt.Println(*p)
	pointer(p)
}