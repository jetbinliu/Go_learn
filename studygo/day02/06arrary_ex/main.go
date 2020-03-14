package main

import "fmt"

func sumArr() {
	arr1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println(sum)
}

func findSum() {
	arr2 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr2); i++ {
		for j := i + 1; j < len(arr2); j++ {
			if arr2[i]+arr2[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}

func main() {
	sumArr()
	findSum()
}
