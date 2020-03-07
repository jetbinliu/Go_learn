//九九乘法表
package main

import (
	"fmt"
)

func multiTable() {
	for m := 1; m <= 9; m++{
		for n := 1; n <=m; n++{
			multi := m * n
			fmt.Printf("%d*%d=%d ", n, m, multi)
		}
		fmt.Println()
	}
}

func main() {
	multiTable()
}