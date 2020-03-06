package main

import (
	"fmt"
	"math/rand"
	//"time"
)

type Student struct {
	Name string
	Age int
	Score float32
	next *Student	//这是一个链表了
}

func trans(p *Student) {
	
	for p != nil {
		fmt.Println(*p)
		p = p.next
		//fmt.Println()
	}
}

func insertTail(p *Student) {
	var tail = p
	for i := 0 ; i < 10 ;i++{
		stu := Student{
			Name : fmt.Sprintf("stu%d",i),
			Age : rand.Intn(20),
			Score : rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
		//fmt.Println(tail)
	}
	
}

func insertHead(head **Student){
	for i := 0; i < 10; i++{
		stu := Student{
			Name : fmt.Sprintf("stu%d",i),
			Age : rand.Intn(20),
			Score : rand.Float32() * 100,
		}
		stu.next = *head
		*head = &stu
		//fmt.Println(*head)
}
}

func delNode(p *Student){
	var prev *Student = p
	for p != nil {
		if p.Name == "stu6" {
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}
}

func addNode(p *Student,newNode *Student) {
	for p != nil {
		if p.Name == "stu0" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
}
}

func main() {
	var head *Student = new(Student)

	head.Name = "one"
	head.Age = 1
	head.Score = 100
	//insertTail(head)
	insertHead(&head)
	trans(head)
	delNode(head)
	trans(head)
	//insertHead(head)
	
	var newNode *Student = new(Student)
	newNode.Name = "stu1000"
	newNode.Age = 18
	newNode.Score = 100
	addNode(head,newNode)
	trans(head)

	// var stu1 Student
	// stu1.Name = "two"
	// stu1.Age = 2
	// stu1.Score = 122

	// var stu2 Student
	// stu2.Name = "three"
	// stu2.Age = 3
	// stu2.Score = 122

	// head.next = &stu1
	// stu1.next = &stu2

	// trans(&head)
	// fmt.Println()
	
	//trans(&head)
	//fmt.Println(head,stu1.next)
}