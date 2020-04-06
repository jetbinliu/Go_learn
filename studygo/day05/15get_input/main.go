package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取输入，无法处理空格
func useScan() {
	var str string
	fmt.Printf("请输入内容:")
	fmt.Scanln(&str)
	fmt.Println("你输入的是:", str)
}

// 获取输入，可以处理空格
func useBufio() {
	var str string
	fmt.Printf("请输入内容:")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read failed,err:", err)
	}
	fmt.Println("你输入的内容是:", str)
}

func main() {
	//useScan()
	useBufio()
}
