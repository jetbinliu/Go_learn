package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json

// 1. 序列化：json.Marshal() 把Go语言中的结构体变量 --> json格式的字符串
// 2. 反序列化：json格式的字符串 --> Go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name"` //在json格式化的时候，字段名字用name代替Name
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "zhangsan",
		Age:  20,
	}
	// 序列化
	b, err := json.Marshal(p1) //对结构体进行调用，结构体中的字段名首字母要大写
	if err != nil {
		fmt.Printf("marshal faild,err:%v", err)
	}
	fmt.Printf("%v\n", string(b))

	// 反序列化
	str := `{"name":"lisi","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在函数内部修改p2的值
	fmt.Println(p2)
}
