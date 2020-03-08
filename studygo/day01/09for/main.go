package main

import "fmt"

func mutiplicationTable1() {
	for i := 9; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func mutiplicationTable2() {
	flag := false
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
			if j == 4 {
				flag = true
				break
			}
		}
		fmt.Println()
		if flag {
			break
		}
	}
}

func test1() {
	for i := 0; i < 10; i++ {

		if i == 5 {
			continue
		}
		fmt.Print(i)
	}
	fmt.Println("over")
}

func main() {
	// mutiplicationTable1()
	mutiplicationTable2()
	//test1()
}
