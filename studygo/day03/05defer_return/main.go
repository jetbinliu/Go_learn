package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
		fmt.Println("bbbb", x)
	}()
	fmt.Println("aaaa", x)
	return x
}

func f2() (x int) {
	defer func() {
		fmt.Println("f2cccc", x)
		x++
		fmt.Println("f2bbbb", x)
	}()
	fmt.Println("f2aaaa", x)
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//fmt.Println(f1())
	fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4())
	//defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	x := 1
	y := 2
	defer calc("aa", x, calc("a", x, y))
	x = 10
	defer calc("bb", x, calc("b", x, y))
	x = 20

}
