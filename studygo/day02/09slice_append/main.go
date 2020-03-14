package main

import "fmt"

//append() 为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "深圳", "上海"}
	fmt.Println(s1, len(s1), cap(s1))
	//s1[3] = "广州" //错误的写法
	//调用append(函数必须用原来的切片变量接收返回值)
	s1 = append(s1, "广州", "杭州")
	fmt.Println(s1, len(s1), cap(s1))

	s2 := []string{"北京", "上海", "深圳", "上海"}
	s1 = append(s1, s2...) //...表示拆开
	fmt.Println(s1, len(s1), cap(s1))

}
