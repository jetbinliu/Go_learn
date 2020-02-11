package main

import "fmt"

func list(n int) {
	for i := 0; i <= n; i++{
		fmt.Printf("%d+%d=%d\n",i, n-i, n)
	}
}

func main() {

	// n := 10
	// for i :=0 ; i <= n; i++{
	// 	fmt.Printf("%d+%d=%d\n",i, n-i, n)
	// }
	list(5)


}