package main

import "fmt"

//函数
//意义：函数是一段代码的封装
//把一段逻辑抽象出来封装到一个函数中，给它起个名字，每次用到它的时候直接用函数名调用就可以了
//使用函数能够让代码结构更清晰、更简洁

//函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数但有返回值
func f3() int {
	return 3
}

//返回值可以命名也可以不命名
//命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return //使用命名返回值可以return后省略
}

//多个返回值
func f5() (int, int) {
	return 1, 3
}

//参数的类型简写
func f6(x, y int, m, n string, i, j bool) int {
	fmt.Println(m, n, i, j)
	return x + y
}

//可变长参数，必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y的类型是切片 []int
}

func main() {
	r := sum(10, 20)
	fmt.Println(r)
	f1(0, 1)
	f2()
	f3 := f3()
	fmt.Println(f3)
	fmt.Println(f4(2, 2))

	m, n := f5()
	fmt.Println(m, n)
	f6 := f6(3, 3, "a", "b", true, false)
	fmt.Println(f6)
	f7("f7", 2, 3, 4, 5)
}
