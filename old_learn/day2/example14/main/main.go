package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

var test string = `hello world \n\n
hello
world
`
var test1 byte = 'c'

func main(){
    
	for i := 0; i < 2; i++ {
		fmt.Println(rand.Int())
		fmt.Println(rand.Intn(100))
		fmt.Println(rand.Float32())
	}
	// fmt.Println(rand.Intn(100))
	// fmt.Println(rand.Float32())
	fmt.Println(test)
	fmt.Println(test1)
	fmt.Printf("%c\n",test1)
	fmt.Printf("%q\n","this is test")
	fmt.Printf("%s\n","this is test")
	str := fmt.Sprintf("a=%d",1)
	fmt.Printf("str=%q , %T\n",str, str)
}