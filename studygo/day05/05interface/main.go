package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chiken struct {
	feet int8
}

func (c chiken) move() {
	fmt.Println("激动")
}
func (c chiken) eat(food string) {
	fmt.Printf("吃%s\n", food)
}

func (c cat) move() {
	fmt.Println("猫步")
}
func (c cat) eat(food string) {
	fmt.Printf("吃%s", food)
}

func main() {
	var a1 animal // 定义一个接口类型的变量
	fmt.Printf("%T\n", a1)
	bc := cat{ // 定义一个cat类型的变量
		name: "淘气",
		feet: 4,
	}
	kfc := chiken{
		feet: 2,
	}
	a1 = bc
	a1.eat("耗子")
	fmt.Printf("%T\n", a1)
	fmt.Println(a1)

	a1 = kfc
	a1.eat("虫子")
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
}
