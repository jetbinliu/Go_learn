package main

import (
	"fmt"
)

func fibnacci(n int) {
	//var a []int
	a := make([]int, n)

	a[0] = 1
	a[1] = 1
	
	for i := 2; i < n ; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for _,v := range a {
		fmt.Println(v)
	}
}

func initArray() {
	var a [5]int = [5]int{1,2,3,4,5}
	var b  = [...]int{1,2,3,4}
	c := [...]int{1}
	d := [...]int{1:100,2:200}

	

	fmt.Println(a,b,c,d)
}

func initArray2() {
	var a [2][5]int = [2][5]int{{1,2,3,4,5},{6,7,8,9,10}}
	for _,v := range a {
		for _,v1 :=range v{
			fmt.Printf("%d ",v1)
		}
		fmt.Println()
	}
}

func main() {
	//fibnacci(10)
	var a []int
	fmt.Println(a)
	initArray2()
}