package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c",str[strLen - 1 - i])
	}
	return result
}

func reverse1(str string) string {

	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result,tmp[length - i -1])
	}
	fmt.Println(string(result))
	return string(result)
}

func main() {

	var str1 = "hello"
	str2 := "world"
	//str3 := str1 + " " + str2
	str3 := fmt.Sprintf("%s %s", str1, str2)
    n := len(str3)

	fmt.Println(str3)
	fmt.Printf("len str3 is %d\n",n)
	
	substr := str3[6:]
	fmt.Println(substr)

	substr1 := strings.Split(str3, " ")
	fmt.Println("split:",substr1[0])

	fmt.Println(reverse1(str3))
	

}