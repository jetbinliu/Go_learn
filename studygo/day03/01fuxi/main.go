package main

import "fmt"

//day03 复习

func main() {
	// var name string
	// name = "理想"
	// fmt.Println(name)
	// var ages [10]int //声明了一个变量，[30]int类型s
	// ages = [10]int{1, 2, 3}
	// fmt.Println(ages)

	// ages2 := [...]int{1, 2, 3, 4, 5, 6}
	// fmt.Println(ages2)

	// var ages3 = [...]int{1: 10, 10: 20}
	// fmt.Println(ages3)

	// //二维数组
	// var a1 = [...][2]int{
	// 	[2]int{1, 2},
	// 	[2]int{3, 4},
	// 	[2]int{5, 6},
	// }
	// fmt.Println(a1)
	// //多维数组只有最外层可以使用...

	// //数组是值类型
	// x := [3]int{1, 2, 3}
	// fmt.Println(x)
	// f1(&x)
	// fmt.Println(x)

	//切片
	// var s1 []int //没有分配内存 ，==nil
	// fmt.Println(s1 == nil)
	// s1 = []int{1, 2, 3}
	// fmt.Println(s1)
	// //make初始化分配内存
	// s2 := make([]bool, 2, 4) //2长度，4容量
	// fmt.Println(s2)
	// s3 := make([]int, 0, 4)
	// fmt.Println(s3 == nil)
	// s1 := []int{1, 2, 3}
	// s2 := s1
	// s2[1] = 200
	// fmt.Println(s2)
	// fmt.Println(s1)
	// var s1 []int
	// s1 = make([]int, 1)
	// s1 = append(s1, 1)
	// fmt.Println(s1)

	// s1 := []int{1, 2, 3}
	// s2 := s1
	// var s3 []int
	// s3 = make([]int, 3)
	// copy(s3, s1)
	// //s3 = append(s3, s1...)
	// fmt.Println(s2)
	// s2[1] = 200
	// fmt.Println(s1)
	// fmt.Println(s2)
	// fmt.Println(s3)

	//指针
	//Go 里面的指针只能读不能修改，不能修改指针变量指向的地址
	// addr := "沙河"
	// addrP := &addr
	// fmt.Println(addrP) //addr变量指向的内存地址
	// fmt.Printf("%T\n", addrP)
	// addrV := *addrP //根据内存地址找值
	// fmt.Printf("%T,%v\n", addrV, addrV)
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 1)
	m1["zhangsan"] = 1
	m1["lisi"] = 2
	m1["wanger"] = 3
	fmt.Println(m1)
	a, ok := m1["jiumging"]
	fmt.Println(a, ok)

}

func f1(a *[3]int) {
	a[1] = 100
	//fmt.Println(a)
}
