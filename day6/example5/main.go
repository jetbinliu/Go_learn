package main

import (
	"fmt"
)

type Student struct {
	Name string
	Sex string
}

func Test(a interface{}) {
	d,ok := a.(int)
	if ok == false {
		fmt.Println("convert failed,type must be int")
		return
	}
	d = 3
	fmt.Println("d:",d)
}

func just(items ...interface{}) {
	for index,v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool,value is %v\n", index,v)
		case int,int64,int32:
			fmt.Printf("%d params is int,value is %d\n", index,v)
		case float32,float64:
			fmt.Printf("%d params is float,value is %f\n", index,v)
		case string:
			fmt.Printf("%d params is float,value is %s\n", index,v)
		case Student:
			fmt.Printf("%d params is student,value is %v\n", index,v)
		case *Student:
			fmt.Printf("%d params is *student,value is %v\n", index,v)
		}
	}
}

func main() {
	var a interface{}
	var b int
	fmt.Printf("%T\n",a)

	a = b
	c := a.(int)
	fmt.Printf("%T\n",c)

	var d int
	Test(d)
	var e int
	Test(e)

	var f Student = Student{
		Name:"zhangsan",
		Sex: "man",
	}
	just("hello world",28,8.2,f,&f)
}