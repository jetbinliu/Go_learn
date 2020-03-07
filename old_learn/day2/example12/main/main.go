package main

import (
	"fmt"
)

var a string

func main(){
	a = "G" //修改了全局变量，后面调用全局变量a的值均为G
	fmt.Println(a)
	f1()
	var b int = 1234567890123456789
	fmt.Println(b)
}

func f1(){
	a := "O"
	fmt.Println(a)
	f2()
}

func f2(){
	fmt.Println(a)
}