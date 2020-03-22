package main

import "fmt"

//结构体模拟实现其他类型语言中的“继承”

type animal struct {
	name string
}

// 给animal实现一个移动的方法

func (a animal) move() {
	fmt.Printf("%s 搜搜跑\n", a.name)
}

type dog struct {
	feet   uint8
	animal //嵌套结构体，animal拥有的方法dog此时也拥有
}

// 给do实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s 汪汪汪~\n", d.name)
}

func main() {
	a := animal{
		name: "dog",
	}
	fmt.Println(a)

	d1 := dog{
		animal: animal{name: "dahei"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.move()
	d1.wang()
}
