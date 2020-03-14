package main

import "fmt"

//切片slice

func main() {
	//切片的定义
	var s1 []int    //定义一个存放int类型的切片
	var s2 []string //string类型切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(len(s1), cap(s1))

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张家口"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	//2.由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	s3 := a1[1:] //基于一个数组切割，左闭右开
	s3[1] = 1
	fmt.Println(s3, len(s3), cap(s3))

	//切片再切片
	s8 := s3[3:]
	a1[5] = 10000
	fmt.Println(s8, len(s8), cap(s8))
}
