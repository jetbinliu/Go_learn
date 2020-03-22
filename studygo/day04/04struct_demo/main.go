package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

func f(x person) {
	(x).gender = "女" //修改的是副本的gender
}

func f2(x *person) {
	//(*x).gender = "女" //根据内存地址找到原变量，修改的就是原来的变量
	x.gender = "女"
}

func main() {
	var p person
	p.name = "zhangsan"
	p.gender = "男"
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)

	//结构体指针1
	var p2 = new(person)
	(*p2).name = "lisi"
	p2.gender = "nan"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  //p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) //求b2的内存地址
	fmt.Println(*p2)
	// 2.1结构体指针,key value初始化
	var p3 = &person{
		name:   "wanger",
		gender: "女",
	}
	fmt.Printf("%#v\n", p3)
	// 2.2 使用值列表的形式初始化，值得顺序要和结构体定义时一致
	p4 := person{
		"mawu", "女",
	}
	fmt.Printf("%v\n", p4)
}
