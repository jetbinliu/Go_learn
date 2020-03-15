package main

import "fmt"

//函数

func f1() {
	fmt.Println("hell world")
}

func f2(name string) {
	fmt.Println("hello", name)
}

func f3(x int, y int) int {
	sum := x + y
	return sum
}

func f5(title string, y ...int) int {
	fmt.Println(y)
	return 1
}

//命名返回值
func f6(x, y int) (sum int) {
	sum = x + y
	a := 1
	return a
}

//go中支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f1()
	f2("a")
	a := f3(100, 200)
	fmt.Println(a)

	f5("lisi", 1, 2, 3, 4, 5)
	fmt.Println(f6(1, 2))
	fmt.Println(f7(1, 2))
}
