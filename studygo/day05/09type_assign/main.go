package main

import "fmt"

// 类型断言
// 类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	_, ok := a.(string)
	if !ok {
		fmt.Println("not string")
	} else {
		fmt.Println("string type")
	}

}

// 类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch a.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case bool:
		fmt.Println("type is bool")
	default:
		fmt.Println("type is other")
	}
}

func main() {
	assign2(11)
	assign2("hello")
	assign2([...]int{1, 2})
}
