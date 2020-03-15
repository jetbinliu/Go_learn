package main

import "fmt"

//defer，用于函数结束之前释放资源

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("hehe")
	fmt.Println("end")
	defer fmt.Println("11")
}
func main() {
	deferDemo()
}
