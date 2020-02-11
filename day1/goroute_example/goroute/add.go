package goroute


func Add(x int, y int, chain chan int){

	sum := x + y
	chain <- sum
}