package main

import (
	"fmt"
)

func testSlice() {
	var a [5]int = [...]int{1,2,3,4,5}
	s  := a[1:]
	fmt.Printf("1len[%d] cap[%d]\n",len(s),cap(s))
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	s= append(s,10)
	s= append(s,10)
	fmt.Printf("2len[%d] cap[%d]\n",len(s),cap(s))
	s= append(s,10)
	s= append(s,10)
	s= append(s,10)
	s= append(s,10)
	fmt.Printf("3len[%d] cap[%d]\n",len(s),cap(s))
	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
}

func testCopy () {
	var a []int = []int{1,2,3,4,5,6}
	b := make([]int,1)
	copy(b,a)
	c := append(b,a...)
	fmt.Println(b)
	fmt.Println(c)
}

func testString() {
	s := "你hello world"
	s1 := s[:5]
	s2 := []rune(s)
	s2[0] = 'a'
	fmt.Println(s1,string(s2))
}

func main() {
	//testSlice()
	//testCopy ()
	testString()
}