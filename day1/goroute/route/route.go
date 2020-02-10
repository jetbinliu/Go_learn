package route

import "fmt"

func Route(chain chan int){

	result := <- chain
	fmt.Println(result)
}