package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
)

func main() {

	str := "   hello world abc world   "
	result := strings.Replace(str, "world", "china",-1)
	fmt.Println(result)

	count := strings.Count(str,"world")
	fmt.Println(count)

	result1 := strings.Repeat(str,3)
	fmt.Println(result1)

	resultToUpper := strings.ToUpper(str)
	fmt.Println(resultToUpper)

	resultToLower := strings.ToLower(str)
	fmt.Println(resultToLower)

	resultTrimeSpace := strings.TrimSpace(str)
	fmt.Println(resultTrimeSpace)

	resultTrime := strings.Trim(str, " \n\r")
	fmt.Println(resultTrime)

	resultTrimeLeft := strings.TrimLeft(str, " \n\r")
	fmt.Println(resultTrimeLeft)

	resultTrimeRight := strings.TrimRight(str, " \n\r")
	fmt.Println(resultTrimeRight)

	resultFields := strings.Fields(str)
	for i := 0; i < len(resultFields); i++{
		fmt.Println(resultFields[i])
	}
	//fmt.Println(resultFields)

	resultSplit := strings.Split(str,"l")
	for i := 0; i < len(resultSplit); i++{
		fmt.Println(resultSplit[i])
	}
	
	resultJoin := strings.Join(resultFields, "22222")
	fmt.Println(resultJoin)

	resultItoa := strconv.Itoa(1000)
	fmt.Printf("%q%T\n",resultItoa,resultItoa)

	resultAtoi,err := strconv.Atoi(string("12345"))
	fmt.Println(resultAtoi,err)

	now := time.Now()
	fmt.Println(now.Format("2006-1-02 15:04"))

}