//求小于n的完数完数（一个数等于它的因子（大于0并小于自己的可以被自己整除的数）之和）
package main

import (
	"fmt"
)

func factor(f int) []int {
	factorArray := []int{}
	for i := 1; i < f; i++ {
		if f%i == 0 {
			factorArray = append(factorArray, i)
			//fmt.Printf("%d ",i)
		}
	}
	return factorArray
}

func perfectNum(n int) []int {
	array := []int{}

	for num := 1; num <= n; num++ {
		sum := 0
		for _, v := range factor(num) {
			sum = sum + v
		}
		if sum == num {
			fmt.Println(sum)
			factorNum := factor(sum)
			fmt.Println(factorNum)
			array = append(array, sum)
		}
	}
	return array
}

func main() {
	a := perfectNum(1000)
	fmt.Println(a)
	// b := factor(6)
	// fmt.Println(b)
}
