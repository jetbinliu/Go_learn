package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}
type workPlace struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	//city    string
	address // 匿名嵌套结构体，address类型的变量
	workPlace
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "zhangsan",
		age:  19,

		address: address{ // 类型名作为变量
			province: "shangdong",
			city:     "qingdao",
		},
		workPlace: workPlace{
			city: "linyi",
		},
		//city: "linyi",
	}
	fmt.Println(p1.workPlace.city) // 先查找自己结构体的字段，找不到采取匿名嵌套的结构体中查找
}
