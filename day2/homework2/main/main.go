package main

import (
	"fmt"
	"strconv"
)

func isFlower(n, m int) {
	for i := n; i < m ; i++{
		var a, b, c int
		a = i % 10
		b = (i / 10) % 10
		c = (i / 100) % 10
		if a*a*a + b*b*b + c*c*c == i {
			fmt.Println(i)
		}
	}
}

func isFlower2(n, m int) {
	for i := n; i < m; i++{
		strI := fmt.Sprintf("%d",i)
		
		var result = 0
		for b := 0; b < len(strI); b++{
			num := int(strI[b] - '0')
			result = result + num*num*num
		}

		number, err := strconv.Atoi(strI)
		if err != nil {
			fmt.Printf("can not convert %s to int\n", strI)
			return
		}

		if result == number {
			fmt.Println(result)
		}
	}
}

func main(){
	//isFlower(100,999)
	isFlower2(1,10000)
}

/*
func main() {

	for a := 100; a <= 999; a++ {
		//b := strconv.Itoa(a)
		b := fmt.Sprintf("%d", a)
		a1, _ := strconv.Atoi(string(b[0]))
		a2, _ := strconv.Atoi(string(b[1]))
		a3, _ := strconv.Atoi(string(b[2]))
		if a1*a1*a1 + a2*a2*a2 + a3*a3*a3 == a{
			fmt.Println(a1, a2, a3, a)
		}
		//fmt.Println(a1, a2, a3)
		// lenb := len(b)
		// for i := 0; i < lenb; i++ {
		// 	str1 := string(b[i])

		// 	fmt.Printf("%q\n",str1)
		// }
	}
}
*/
