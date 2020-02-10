package prt_fmt

import "fmt"

func Char_fmt() {
	x := "a"
	y := 17
	fmt.Printf("x is %s, y is %d\n", x, y)
	fmt.Printf("x B is %b, y B is %b\n", x, y)
	fmt.Printf("x O is %o, y O is %o\n", x, y)
	fmt.Printf("x D is %d, y D is %d\n", x, y)
	fmt.Printf("x H is %x, y H is %x\n", x, y)
	fmt.Printf("x f is %f, y f is %f\n", x, y)

}