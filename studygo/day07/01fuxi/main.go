package main

import (
	"encoding/json"
	"fmt"
)

// 反射
type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"zhangsan","age":16}`
	fmt.Printf("type:%T,value:%v\n", str, str)
	var stu1 student
	json.Unmarshal([]byte(str), &stu1)
	//fmt.Printf("%#v\n", stu1)
	fmt.Printf("type:%T,value:%v\n", stu1, stu1)
}
