package main

import "fmt"

// 引出接口的实例
type speaker interface {
	speak() //方法的签名
}

type cat struct{}
type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}
func (d dog) speak() {
	fmt.Println("旺旺旺~")
}

func da(x speaker) {
	//接收一个参数，传进来什么就打什么
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)

}
