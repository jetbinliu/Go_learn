package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a int = rand.Intn(100)
	

	for {
		var input int
		fmt.Scanf("%d",&input)
        flag := false
	    switch  {
		case input == a:
			fmt.Println("U right")
			flag = true
		case input > a:
			fmt.Println("big")
			default :
			fmt.Println("small")
		}

		if flag == true{
			break
		}
	}

	 

    
	
	
	/*
	switch a {
	case 0,1,2,3:
		fmt.Println("a is 0")
		fallthrough
	case 10:
		fmt.Println("a is 10")
	default:
		fmt.Println("other")
	}

	switch {
	case a > 0 && a < 10 :
		fmt.Println("a > 0 and a < 10")
	default:
		fmt.Println("aaaa")
	}

	switch b := 30; {
	case b > 0 && b <10:
		fmt.Println("b>0 and b< 10")
	default:
		fmt.Println("bbbb")
	}

	*/
}
