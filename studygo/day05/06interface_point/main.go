package main

import "fmt"

// 使用值接收者和指针接收者的区别
type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}

// 使用值接收者
// func (c cat) move() {
// 	fmt.Println("猫步")
// }
// func (c cat) eat(food string) {
// 	fmt.Printf("猫吃%s\n", food)
// }

// 使用指针接收者
func (c *cat) move() {
	fmt.Println("猫步")
}
func (c *cat) eat(food string) {
	fmt.Printf("吃%s", food)
}

func main() {
	var a1 animal
	c1 := cat{"miaomiao", 4} // cat
	c2 := &cat{"tom", 4}     // *cat

	a1 = &c1
	//a1.move()
	a1.eat("鱼")
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
