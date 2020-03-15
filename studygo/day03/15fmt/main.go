package main

import "fmt"

func main() {
	//获取用户输入
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("type:", s)

	var (
		name  string
		age   int
		clase string
	)
	// fmt.Scanf("%s %d %s\n", &name, &age, &clase)
	// fmt.Println(name, age, clase)
	fmt.Scanln(&name, &age, &clase)
	fmt.Println(name, age, clase)

}
