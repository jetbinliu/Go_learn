package main

import (
	"fmt"
)


func test1() {
	var a [10]int
	//j := 10
	a[0] = 100
	a[9] = 200
	fmt.Println(a)

	for i:=0; i< len(a);i++{
		fmt.Println(a[i])
	}

	for _,v := range a {
		fmt.Println(v)
	}
}

func test2() {
	var a [10]int
	b := a

	b[0] = 100
	fmt.Println(a)
}

func test3(arr *[5]int) {
	(*arr)[0] = 1000
}


func main() {
	//test1()
	test2()
	
	var a [5]int
	test3(&a)
	fmt.Println(a)
}