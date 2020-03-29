package main

import (
	"encoding/json"
	"fmt"
)

// 复习结构体

type tmp struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

func newPerson(n string, i int) (p person) {
	p = person{
		name: n,
		age:  i,
	}
	return p
}

//对函数person造方法
//接收者使用对应类型的首字母小写
//指定了接收者，只有接收者这个类型的变量才能调用这个方法

func (p person) dream(str string) {
	fmt.Printf("%s的梦想是%s\n", p.name, str)
}

//指针接受者
//1.需要修改结构体变量的值时要使用指针接受者
//2.结构体本身比较大，拷贝的内容开销比较大时也要使用指针接收者
//3.保持一致性：如果一个方法使用了指针接收者，其他的方法为了统一也要使用指针接收者
func (p *person) guonian() {
	p.age++
}

func main() {

	type point struct {
		X int `json:"x"` //X需要大写，不然后面无法调用，``表示json调用时显示为x
		Y int `json:"y"`
	}
	b1 := point{100, 200}
	c, err := json.Marshal(b1) // json序列化，GO通常将第二个参数的返回值作为error
	if err != nil {
		fmt.Println("序列化出错了:", err)
	}
	fmt.Println(string(c)) //Marshal返回的是[]byte

	// 反序列化：由字符串 --> GO 中的结构体变量
	str1 := `{"x":6600,"y":1100}`
	var po2 point                      // 造一个结构体变量，准备接受反序列化的值
	json.Unmarshal([]byte(str1), &po2) // 传入的需要时byteLeixing ，结构体变量需为指针
	if err != nil {
		fmt.Println("反序列化出错了:", err)
	}
	fmt.Println(po2)

	// 结构体的嵌套
	type addr struct {
		province string
		city     string
	}
	type student struct {
		name    string
		address addr //嵌套的结构体,addr为类型，不写名称为匿名嵌套：适用类型做名称
	}

	var p1 = person{"zhangsan", 19}
	(&p1).guonian()
	p1.guonian() //可以不用&，直接简写
	fmt.Println(p1)

	var b = tmp{
		20,
		10,
	}
	fmt.Println(b)
	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)
}
