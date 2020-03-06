package main

import (
	"fmt"
)

type integer int
func (p integer) print(){
	fmt.Println("p is",p)
}
func (p *integer) set(b integer) {
	*p = b
}

type Student struct {
	Name string
	Age int
	Score int
}

func (p *Student) init(name string,age int,score int) {
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(*p)
}

func (p Student) get() Student{
	return p
}

func main() {
	var stu Student
	stu.init("stu",10,200)
	//fmt.Println(stu)
	stu1 := stu.get()
	fmt.Println(stu1)

	var a integer = 100
	a.print()
	a.set(200)
	a.print()
}