package main

import (
	"fmt"
	"sort"
)

func testMapSort() {
	a := make(map[int]int,10)
	a[1] = 10
	a[33] = 10
	a[53] = 10
	a[23] = 10
	a[44] = 10

	var keys []int
	for k,_ := range a{
		keys = append(keys,k)
	}
	sort.Ints(keys)

	for _,v := range keys{
		fmt.Println(v,a[v])
	}

	fmt.Println(a)
}

func testReverse() {
	a := make(map[string]int,10)
	b := make(map[int]string,10)
	a["abc"] = 120
	a["e"] = 220
	fmt.Println(a)
	for k,v := range a{
		b[v] = k
	}
	fmt.Println(b)

}

func main() {
	//testMapSort()
	testReverse()
}