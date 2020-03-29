package main

import "fmt"

//接口2
//不管什么牌子的车，都能跑进行

type car interface {
	run()
}

type falali struct {
	brand string
}

func (f falali) run() {
	fmt.Println("falali sousousou")
}

type baoshijie struct {
	brand string
}

func (b baoshijie) run() {
	fmt.Println("baoshijie wengwengweng")
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)
}
