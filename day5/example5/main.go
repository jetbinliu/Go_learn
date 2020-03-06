package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score int `json:"score"`
}

func main() {
	var stu Student = Student {
		Name:"zhangsan",
		Age:11,
		Score:100,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu faild,err:",err)
	}
	fmt.Println(string(data))
	
}