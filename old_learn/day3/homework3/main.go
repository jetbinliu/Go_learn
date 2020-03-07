package main

import (
	"fmt"
)

func overturn(str string) string{
	var tmpStr string
	for i := len(str)-1; i >= 0; i-- {
		tmpStr = tmpStr + string(str[i])
	}
	return tmpStr
}

func main() {
	var a string
	fmt.Scanf("%s",&a)
	b := overturn(a)
	fmt.Printf("a=%s, b=%s\n",a, b)

	if b == a {
		fmt.Println("It is plalindrome.")
	} else {
		fmt.Println("It is not plalindrome!")
	}
}
