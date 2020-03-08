package main

import "fmt"

const pi = 3.1415

const (
	a1 = 100

	a2
	a3
)

const (
	n1 = iota + 100
	n2
	n3 = 30
	_
	n4
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //表示1向左移动10*1=10位，b以二进制表示，转换为十进制为1024
	MB = 1 << (10 * iota) //表示1向左移动10*2=20位，b以二进制表示，转换为十进制为1,048,576
	GB = 1 << (10 * iota) //表示1向左移动10*3=30位，b以二进制表示，转换为十进制为1,073,741,824
	TB = 1 << (10 * iota) //表示1向左移动10*4=40位，b以二进制表示，转换为十进制为1,099,511,627,776
)

func main() {
	//pi = 123
	fmt.Println(n1, n2, n3, n4)
	fmt.Println(a1, a2, a3)
	fmt.Println(KB, MB, GB, TB)
}
