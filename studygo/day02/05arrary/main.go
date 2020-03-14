package main

import "fmt"

//数组
//存放元素的容器
//必须制定存放的元素的类型和容量（长度）
//数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T,a2:%T\n", a1, a2)

	//数组的初始化
	fmt.Println(a1, a2)
	//1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//2.初始化方式2（长度以...表示）
	a100 := [...]int{1, 2, 3, 4, 2, 3, 4}
	fmt.Println(a100)
	//3.初始化方式3（根据索引初始化）
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2.for range
	for k, v := range citys {
		fmt.Println(k, v)
	}

	//多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{4, 5},
	}
	fmt.Println(a11)
	//多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[1] = 100
	fmt.Println(b1, b2)
}
