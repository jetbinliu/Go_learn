package main

import "fmt"

//关于append删除切片中的某个元素
func main() {
	a1 := [...]int{1, 3, 5, 7, 2, 4, 15, 14}
	s1 := a1[:]

	//删掉索引为1的3
	s1 = append(s1[0:1], s1[2:]...)

	a1[0] = 200
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(a1)

}
