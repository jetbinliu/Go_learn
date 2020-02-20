package main

import (
	"fmt"
)

func add(a int, b...int) (sum int){
	sum = a
	for _,v := range b {
		sum = sum + v
	}
	//sum = a + sumb
	return

}

func main() {
	c := add(1,2,3,4)
	fmt.Println(c)
}