package main

import "fmt"

//构造函数

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := *newPerson("zhangsan", 18)
	p2 := *newPerson("lisi", 20)
	fmt.Println(p1, p2)
	d1 := newDog("heibei")
	fmt.Println(d1)
}
