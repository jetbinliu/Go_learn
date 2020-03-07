package main

import (
	"fmt"
)

type Carer interface{	//接口
	GetName() string	//方法
	Run()	
	DiDi()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string{
	fmt.Println(p.Name)
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running\n",p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s is Didi\n",p.Name)
}

type BYD struct {
	Name string
}

func (p *BYD) GetName() string{
	fmt.Println(p.Name)
	return p.Name
}

func (p *BYD) Run() {
	fmt.Printf("%s is running\n",p.Name)
}

func (p *BYD) DiDi() {
	fmt.Printf("%s is Didi\n",p.Name)
}

func main() {
	var car Carer
	// fmt.Println(car)

	var bmw BMW
	bmw.Name = "baoma"
	car = &bmw
	car.GetName()
	car.Run()

	// var byd BYD
	// byd.Name = "byd"
	// car = &byd 
	byd := &BYD {
		Name : "byd",
	}
	car = byd
	car.GetName()
	car.DiDi()
	//fmt.Println(car.)
}