package main

import "fmt"

// 关闭通道

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	// for x := range ch1 {
	// 	fmt.Println(x)
	// }
	// <-ch1
	// <-ch1
	// y, ok := <-ch1
	// fmt.Println(y, ok)

	for {
		z, ok := <-ch1
		if !ok {
			break
		} else {
			fmt.Println(z)
		}

	}

}