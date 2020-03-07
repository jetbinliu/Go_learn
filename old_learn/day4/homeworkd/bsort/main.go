package main

import (
	"fmt"
)


func bsort(a []int) {
	for i := 0; i < len(a);i++{
		for j := 1; j < len(a)-i; j++{
			if (a[j] < a[j-1]){		//从小到大排序
				a[j],a[j-1] = a[j-1],a[j]
			}
		}
	}
}

func main() {
	b := [...]int{8,7,10,1,56,33}
	bsort(b[:])
	fmt.Println(b)
}