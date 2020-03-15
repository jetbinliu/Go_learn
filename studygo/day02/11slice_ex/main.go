package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]int, 5, 10)
	fmt.Println(a, len(a), cap(a)) //迷惑人的地方
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a, len(a), cap(a))

	a1 := [...]int{4, 1, 9, 7, 3}
	fmt.Printf("%T\n", a1)
	sort.Ints(a1[:]) //对切片进行排序
	fmt.Println(a1)
}
