package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	score float32
}

func main(){
	var stu Student
	stu.Name = "zhangsan"
	stu.Name = "zs"
	stu.Age = 12

	var stu1 *Student = &Student{
		Age:20,
		Name:"lisi",
	}
	var stu2 = Student{
		Age:11,
		Name:"wanger",
	}

	fmt.Println(stu.Name)
	fmt.Printf("Name:%p\n",&stu.Name)
	fmt.Printf("Age:%p\n",&stu.Age)
	fmt.Printf("score:%p\n",&stu.score)
	fmt.Println(*stu1)
	fmt.Println(stu2)
}