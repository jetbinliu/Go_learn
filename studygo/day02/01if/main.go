package main

import "fmt"

func testif() {
	n := 1
	if n > 0 {
		fmt.Println("11")
	} else if n >= 1 {
		fmt.Println("12")
	}
}

func main() {
	testif()
}
