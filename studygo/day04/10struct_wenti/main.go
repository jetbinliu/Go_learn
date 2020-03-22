package main

import "fmt"

//结构体遇到的问题

// 1.自定义类型，myInt(100)是个啥
type myInt int

func (m myInt) hello() {
	fmt.Println("i am int")
}

// 2.结构体初始化
type person struct {
	name string
	age  int
}

// 3.为什么要用构造函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	// 定义变量
	// m := myInt(10)
	// m.hello()
	// fmt.Println(m)

	// 结构体初始化
	// 方法1 一般初始化
	// var p person
	// p.name = "张三"
	// p.age = 18
	// fmt.Println(p)
	// var p1 person
	// p1.name = "李四"
	// p1.age = 20
	// fmt.Println(p1)
	// //方法2 键值对初始化
	// s1 := []int{1, 2, 3}
	// s2 := map[string]int{
	// 	"stu1": 10,
	// 	"stu2": 20,
	// }
	// fmt.Println(s1, s2)
	// var p2 = person{
	// 	name: "网二",
	// 	age:  21,
	// }
	// fmt.Println(p2)
	// // 值类型初始化
	// var p3 = person{
	// 	"马武",
	// 	100,
	// }
	// fmt.Println(p3)

	// 构造函数初始化，使用
	x := newPerson("lisi", 20)
	fmt.Println(x.age)

}
