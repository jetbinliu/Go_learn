package main

import "fmt"

// 空接口

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// interface:关键字
// interface{} :空接口类型
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16) // 使用空接口可以对map存入各种类型的值
	m1["name"] = "周林"
	m1["age"] = 90
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "run"}
	fmt.Println(m1)

	show(false)
	show(m1)
}
