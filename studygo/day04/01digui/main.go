package main

import "fmt"

//递归：函数自己调用自己
//技术n的阶乘
func f(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

//台阶
func taijie(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	ret := f(5)
	fmt.Println(ret)
	tj := taijie(4)
	fmt.Println(tj)
}
