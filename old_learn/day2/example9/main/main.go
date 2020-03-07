package main

import (
	"fmt"
)

func swap(x *int, y *int ) {
	
	temp := *x
	*x = *y
	*y = temp
}

func swap2(x, y int) (int ,int) {
	return y ,x 
}

func main() {

	a := 1
	b := 2
	fmt.Printf("org:a=%d, b=%d\n", a, b)

	// swap(&a, &b)
	// fmt.Printf("swap:a=%d, b=%d\n", a, b)

	// a,b = swap2(a, b)
	// fmt.Printf("swap2:a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("swap3:a=%d, b=%d\n", a, b)
}