package main

import (
	"fmt"
)

type Car struct {
	weight int
	name string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	wheel int
}

type Train struct {
	Car
}

func (p *Train) String() string{
	str := fmt.Sprintf("name:[%s] weight=[%d]",p.name,p.weight)
	return str
}

func main() {
	var a Bike
	a.weight = 100
	a.name = "bike"
	a.wheel = 2
	fmt.Println(a)
	a.Run()

	var b Train
	b.name = "train"
	fmt.Printf("%s\n",&b)
	b.Run()
}