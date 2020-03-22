package main

import "fmt"

//结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一percon类型的变量
	var stu1 person
	//通过字段复制
	stu1.name = "zhangsan"
	stu1.age = 10
	stu1.gender = "男"
	stu1.hobby = []string{"篮球", "足球"}
	fmt.Println(stu1)
	fmt.Println(stu1.name)

	//匿名结构体，函数内部临时用
	var s struct {
		name string
		age  int
	}

	s.name = "lisi"
	s.age = 14
	fmt.Println(s)

}
