package main

import "fmt"

func main() {
	var a interface{} // 定义一个空接口变量a

	a = 100
	v, ok := a.(int8)
	if ok {
		fmt.Println("猜对了，a是int8", v)
		return
	}

	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int64:
		fmt.Println("int64", v2)
	}
}
