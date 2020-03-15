package main

import (
	"fmt"
	"strings"
)

//传入字符串，计算字符串中单词的出现的频率
func countWord(str string) {
	fmt.Println(str)

	sliceWorld := strings.Split(str, " ")
	var sliceTmp = make(map[string]int, 0)

	for _, v := range sliceWorld {
		//count:=0

		if _, ok := sliceTmp[v]; ok {
			sliceTmp[v]++
		} else {
			sliceTmp[v] = 1
		}
	}
	fmt.Println(sliceTmp)
	//var strTmp []string
	//strTmp = make([]string, 10)
	// for k, v := range str {
	// 	fmt.Println(k, string(v))
	// 	if string(v) == " " {
	// 		continue
	// 	}
	// 	// strTmp[i] = strings.join(str[i])
	// }
}

func main() {
	countWord("how do you do haha ha ha ha")
	// type Map map[string][]int
	// m := make(Map)
	// s := []int{1, 2}
	// s = append(s, 3)
	// fmt.Printf("%+v\n", s)
	// m["q1mi"] = s
	// s = append(s[:1], s[2:]...)
	// fmt.Printf("%+v\n", s)
	// fmt.Printf("%+v\n", m["q1mi"])

}
