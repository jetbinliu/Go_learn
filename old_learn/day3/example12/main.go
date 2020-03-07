package main

import (
	"fmt"
	//"strings"
)

func modify (a int) {
	a = 10
}

func modify1 (a *int) {
	*a = 10
}

func main() {
	a := 8
	fmt.Println(a)
	modify(a)
	fmt.Println(a)

	//b := &a
	modify1(&a)
	fmt.Println(a)

}



/*
func main() {
	// for i := 0; i < 5; i++ {
	// 	strRepeat := strings.Repeat("a", i+1)
	// 	fmt.Println(strRepeat)
	// }
	// var str string = "hello world 中国"
	// for k, v := range str {
	// 	if k > 2 {
	// 		fmt.Println("a",string(v))
	// 		continue
	// 	}
	// 	if k > 3{
	// 		fmt.Println("b",v)
	// 		break

	// 	}
	// }

	// LABLE1:
	// for i := 0 ; i < 5; i++{
	// 	for j := 0; j<5; j++{
	// 		if j == 3{
	// 			continue LABLE1
	// 		}
	// 		fmt.Printf("i is :%d, and j is : %d\n",i,j)
	// 	}
	// }
    // i := 0
	// HERE:
	// fmt.Println(i)
	// i++
	// if i == 6{
	// 	return
	// }
	// goto HERE

	// i := 0
	// for {
	// 	if i >= 3{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++;
	// }

	// for i := 0; i < 7; i++{
	// 	if i % 2 == 0{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	
}

func add(a, b int) int{
	return a+b
}

func main() {
	c := add
	fmt.Println(c)

	sum := c(10,20)
	fmt.Println(sum)
	if ( c == add ) {
		fmt.Println("c equal add")
	}
}



type op_func func(int, int) int

func add(a, b int) int{
	return a+b
}

func operator(op op_func, a , b int) int {
//func operator(op func(int, int) int, a , b int) int {
	return op(a, b)
}

func main() {
	c := add
	fmt.Println(c)
	sum := operator(c,100, 200)
	fmt.Println(sum)
}
*/