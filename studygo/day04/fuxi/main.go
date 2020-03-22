package main

import "fmt"

func f1(name string) {
	fmt.Println("hello", name)
}

func f2(f func(string), name string) {
	f(name)
}

func f3() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func low(f func()) {
	f()
}

//闭包
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main() {
	f2(f1, "world")
	a := f3()
	b := a(1, 2)
	fmt.Println(b)

	//闭包调用
	fc := bi(f1, "china")
	low(fc)
}
