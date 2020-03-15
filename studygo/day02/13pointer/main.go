package main

import "fmt"

//vscode 不支持go moduel
//pointer

func main() {
	//1.&：取地址

	// n := 18
	// fmt.Println(&n)
	// p := &n
	// fmt.Printf("%T\n", p) // *int int类型的指针
	// //2. *：根据地址取值，值内存的是地址
	// m := *p
	// fmt.Println(m, p)

	var a = new(int) //指针的声明方法
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)
	var b *int
	fmt.Println(b)

}
