package main

import "fmt"

func main() {
	//fmt.Println("hello")
	//1.判断字符串中汉字的数量
	// s1 := "hello沙河有沙˚˚˚∆∆"
	// var count int
	// for _, v := range s1 {
	// 	if unicode.Is(unicode.Han, v) {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	//2. how do you do中单词出现的次数
	// s2 := "how do you do"
	// s3 := strings.Split(s2, " ")
	// m1 := make(map[string]int, 10)
	// for _, w := range s3 {
	// 	if _, ok := m1[w]; !ok {
	// 		m1[w] = 1
	// 	} else {
	// 		m1[w]++
	// 	}
	// }
	// fmt.Println(m1)

	//3.回文判断（回文：字符串从左往右读和从右往左读是一样的）
	ss := "a上海海上"

	//字符串字符放入一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println(r)
	for i := 0; i < len(r)/2; i++ {
		fmt.Println(i, r[i])
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

}
