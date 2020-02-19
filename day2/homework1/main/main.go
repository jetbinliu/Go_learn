package main

import (
	"fmt"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++{
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var m int
	//fmt.Scanf("%d%d", n,m)
	fmt.Println(n,m)

	for i := n; i < m; i++{
		if isPrime(i) == true{
			fmt.Printf("%d\n", i)
			continue
		}
	}
}


/*
func isPrime(n , m int) {
	for i := n; i < m; i++{
		var flag bool = true
		for a := 2; a < i; a++{
			if (i % a == 0 ) {
				flag = false
				break
			}
		}

		if flag == true {
			fmt.Println(i)
		}
	}
}
func main() {
	var n int
	var m int
	fmt.Scanf("%d%d", &n, &m)
	isPrime(n,m)

}
*/



// func main() {
// 	for i := 100; i < 200 ;i++{
// 		var flag bool = true
// 		for a := 2; a < i; a++{
// 			if (i % a == 0) {
// 				flag = false
// 				break
// 			} 
// 		}
// 		if flag == true {
// 			fmt.Println(i)
// 		}
// 	}
// }




// func main() {

// 	var n int
// 	fmt.Scanf("%d",&n)

// 	var flag bool = true

// 	for i := 2; i < n; i++{
// 		if n % i == 0{
// 			flag = false
// 			break
// 		}
// 	}

// 	if flag == true {
// 		fmt.Printf("n[%d] shi sushu\n",n)
// 	}
// }