package main

import "fmt"

//变量的作用域

var x = 100

func f1() {
	//函数中查找变量的顺序
	//1.先在函数内部查找
	//2.内部没有，查找全局变量
	x := 200
	y := 10
	fmt.Println(x, y)
	if x == 200 {
		fmt.Println(true)
	}
}

func main() {

	f1()
	fmt.Println(x)
}
