package main

import "fmt"

//自定义类型和类型别名

//type后面跟的是类型
type myInt int     //自定义类型
type yourInt = int //类型别名

func main() {
	var a myInt = 1
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b yourInt = 2
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c rune
	c = '中'
	fmt.Printf("%T\n", c)
}
