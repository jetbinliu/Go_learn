package main

import (
	"fmt"
)

func ssort(a []int) {

	for i := 0; i < len(a); i++ {
		var min int = i
		for j := i + 1; j < len(a) ; j++{
			if a[min] > a[j]{
				min = j
				
			}
		}
		a[i],a[min] = a[min],a[i]
	}
}

func main() {
	b := [...]int{3,1,6,11,5}
	ssort(b[:])
	fmt.Println(b)
}