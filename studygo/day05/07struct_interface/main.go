package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口可以嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	age  int8
}

// cat实现了mover接口
func (c *cat) move() {
	fmt.Println("猫步...")
}

// cat实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 mover
	c1 := cat{
		"taoqi",
		18,
	}
	a1 = &c1
	//a1.eat("鱼")
	a1.move()
	fmt.Println(a1)

}
