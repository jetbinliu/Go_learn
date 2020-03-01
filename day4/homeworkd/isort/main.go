package main

import (
	"fmt"
)

func isort(a []int) {
	for i := 1; i < len(a); i++{
		for j := i; j > 0; j--{
			if a[j] < a[j-1]{
				a[j],a[j-1] = a[j-1],a[j]
			}
		}
	}
}

func main() {
	b := [...]int{5,3,1,88,55,14,15}
	isort(b[:])
	fmt.Println(b)
}