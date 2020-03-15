package main

import "fmt"

//闭包
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定一个函数对f2进行包装，f3返回值类型和f1一致
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		//fmt.Println(x + y)
		f(x, y)
	}
	return tmp
}

func main() {
	//把f2作为参数传入f1
	ret := f3(f2, 100, 200) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数,f3靠
	f1(ret)

	// ret := lixiang(f2, 100, 200)
	// ret()
}

// func lixiang(x func(int, int), m, n int) func() {
// 	tmp := func() {
// 		x(m, n)
// 	}
// 	return tmp
// }
