package main

import "fmt"

//copy

func main() {
	// a1 := []int{1, 2, 3}
	// a2 := a1
	// var a3 = make([]int, 3)
	// copy(a3, a1)
	// fmt.Println(a1, a2, a3)
	// a1[0] = 100
	// fmt.Println(a1, a2, a3)

	// //删除a1中的2
	// a1 = append(a1[:1], a1[2:]...)
	// fmt.Println(a1, len(a1), cap(a1))

	s1 := [...]int{1, 3, 5}
	s2 := s1[:]
	s2 = append(s2[:1], s2[2:]...)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s1, len(s1), cap(s1))
}
