package main

import "fmt"

//匿名函数（一般用于函数内部）
var f1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	f1(10, 20)
	//函数内部无法声明带名字的函数，所以定义匿名函数
	f2 := func(x, y int) {
		fmt.Println(x + y)
	}
	f2(1, 2)
	//只调用一次，可以定义写成立即执行函数的匿名函数
	func(x, y int) {
		fmt.Println(x + y)
	}(3, 4)
}
