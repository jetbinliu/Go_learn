package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	a := [...]int{1,33,5,6,2,1,2,8}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testStrSort() {
	a := [...]string{"abc","efg","acb","cda"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloatSort() {
	a := [...]float64{2.3,0.8,1.1,9.8,7.6}
	sort.Float64s(a[:])
	fmt.Println(a)
}

func testIntSearch() {
	a := [...]int{1,7,3,9,2}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:],3)
	fmt.Println(index)
}

func testStringSearch() {
	a := [...]string{"c","d","a","ddf","d"}
	sort.Strings(a[:])
	fmt.Println(a)
	index := sort.SearchStrings(a[:],"a")
	fmt.Println(index)
}

func main() {
	// testIntSort()
	// testStrSort()
	// testFloatSort()
	// testIntSearch()
	testStringSearch()
}