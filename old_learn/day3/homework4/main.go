//输入字符，统计英文字符、空格、数字和其他字符的个数
package main

import (
	"os"
	"bufio"
	"fmt"
)

func count(str string) (worldCount, spaceCount,numberCount,otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >='a' && v <= 'z' :
			fallthrough
		case v >='A' && v <= 'Z' :
			worldCount++
		case v == '_':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result,_ ,err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:",err)
		return
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Println(wc,sc,nc,oc)
}

/*
import (
	"fmt"
	"regexp"
	//"unicode"
)
func countInput(str string) (countStr, countNum, countSpace, countOther int) {
	// var countStr int
	// var countNum, countSpace, countOther int
	//s1:="abcdefghi"
	//for i := 0; i < len(str); i++ {
	for _, tmpStr := range str {
		//tmpStr := string(str[i])
		//fmt.Println(tmpStr)
		matchChar, _ := regexp.Match("[a-zA-Z]", []byte(string(tmpStr)))
		//fmt.Println(matchChar)
		matchNum, _ := regexp.Match("[0-9]", []byte(string(tmpStr)))
		matchSpace, _ := regexp.Match("_", []byte(string(tmpStr)))
		//fmt.Println(matchSpace)

		if matchSpace  {
			//fmt.Println(matchChar)
			countSpace = countSpace + 1
	        //continue
		} else if matchNum {
			countNum = countNum + 1
		} else if matchChar {
			//fmt.Println(matchSpace)
			countStr = countStr + 1
		} else {
			countOther = countOther + len([]rune{tmpStr})
		}
	}
	return //countStr, countNum, countSpace, countOther
}
func main() {
	var input string
	fmt.Scanf("%s", &input)
	a, b, c, d := countInput(input)
	fmt.Printf("count char is:%d , count num is %d , count _ is %d , count other is %d\n",a, b, c, d)
	//fmt.Printf("%T\n", input)
	//matchChar, _ := regexp.Match("[a-zA-z]", []byte(input))
	//fmt.Println(matchChar)
	//e := unicode.Is(unicode.Scripts["Han"], rune("cba"))
	//fmt.Println(e)
}
*/
