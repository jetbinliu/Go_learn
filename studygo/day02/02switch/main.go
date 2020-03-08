package main

import "fmt"

func testSwtich() {
	n := 3
	switch {
	case n > 1:
		fmt.Println("31")
		fallthrough
	case n < 2:
		fmt.Println("32")
	}
}

func main() {
	testSwtich()
}
