package main

import (
	"fmt"
)

func testSlice() {
	var slice []int
	var arr [5]int = [...]int{1,2,3,4,5}
	slice = arr[2:5]
	fmt.Println(slice) 
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))



}

func testSlice4() {
	var a = [10]int{1,2,3,4}
	b := a[0:5]
	fmt.Printf("%d,%p\n",b,b)
	fmt.Printf("%d,%p\n",a[0],&a[0])
}

func main() {
	//fmt.Println(testSlice())

	testSlice4()
}