package main

import "fmt"

func traversalString() {
	s := "hello世界"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, v := range s {
		//fmt.Printf("%v(%c)", s[k], s[k])
		fmt.Printf("%v(%c)", v, v)
	}
}

func changeString() {
	a := 0
	s1 := "hello世界哈哈"
	for _, v := range s1 {
		fmt.Println(string(v), len(string(v)))
		if len(string(v)) == 3 {
			a += 1
		}
	}
	fmt.Println(a)
	// s2 := []rune(s1)
	// s2[0] = '你'
	// for _, v := range s2 {
	// 	c := fmt.Sprint(v)
	// 	fmt.Println(c)
	// }
	// d := fmt.Sprintln(s2)
	// fmt.Println(d)
	// fmt.Println(string(s2))
}

func main() {
	//traversalString()
	//changeString()
	/*
		str1 := "hello world"
		str2 := "你好"
		var str3 rune = '你'
		catalog := "/opt/src/main/go"

		c := strings.Split(catalog, "/")
		fmt.Println(c)

		d := strings.Join(c, "+")
		fmt.Println(d)
		// fmt.Println(strings.HasPrefix(str1, "h"))
		// fmt.Println(strings.HasSuffix(str1, "he"))
		// fmt.Println(strings.Index(str1, "w"))
		// fmt.Println(len(str1), len(catalog), c[0], len(d))
		fmt.Println(len(str1), len(str2), str3)
	*/
}
