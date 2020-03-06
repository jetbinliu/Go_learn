package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score float32
	left *Student
	right *Student
}

func trans(root *Student) {
	if (root == nil) {
		return
	}

	//fmt.Println(root)
	trans(root.left)
	//fmt.Println(root)
	trans(root.right)
	fmt.Println(root)
}


func main(){

	var root *Student = new(Student)
	root.Name = "root"
	root.Age = 13
	root.Score = 100

	var left1 *Student = new(Student)
	left1.Name = "left1"
	left1.Age = 15
	left1.Score = 99
	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "right1"
	right1.Age = 15
	right1.Score = 99
	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "left2"
	left2.Age = 15
	left2.Score = 99
	left1.left = left2

	//fmt.Println(root)
	
	trans(root)
}