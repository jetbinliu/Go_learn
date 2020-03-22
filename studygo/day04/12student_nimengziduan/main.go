package main

import "fmt"

//匿名字段
//字段比较少，比较简单，
//不常用
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"zhangsan",
		10,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)

}
