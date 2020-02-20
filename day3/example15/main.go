package main

import (
	"fmt"
	//"strings"
)

func concat( a string, b...string) (c string) {
	var conb string
	for _,v := range b {
		conb = conb + v
	}
	c = a + conb
	return
}

func main() {
	d := concat("a","b"," c","nihao")
	fmt.Println(d)
}