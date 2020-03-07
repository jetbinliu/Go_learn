package main

import (
	"fmt"
)

type Test interface{
	Print()
}

type Student struct {
	name string
	age int
	score int
}

func (p *Student) Print() {
	fmt.Println("name:",p.name)
	fmt.Println("age:",p.age)
	fmt.Println("score:",p.score)
}

func main() {

	var t Test
	
	var p Student
	p.name = "zhsngas"
	p.Print()
	t = &p
	t.Print()
}