package main

import (
	"fmt"
)

type integer int

func main() {
	var i integer = 1000
	var j int = 100
	j = int(i)
	//i = 1
	fmt.Println(i, j)
}