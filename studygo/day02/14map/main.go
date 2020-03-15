package main

import "fmt"

//map
func main() {
	var m1 map[string]int = make(map[string]int, 10) //必须先初始化（内存中开辟空间）,10为容量
	m1["李想"] = 1
	m1["zhangsan"] = 2
	fmt.Println(m1)
	// fmt.Println(m1["zhangsn"])
	// value, ok := m1["lisi"]
	// if !ok {
	// 	fmt.Println("no")
	// } else {
	// 	fmt.Println(value)
	// }
	// c := m1["wanger"]
	// fmt.Printf("%T\n", c)

	for k, v := range m1 {
		fmt.Println(k, v)
	}
	delete(m1, "zhangsan")
	fmt.Println(m1)
}
