package main

import (
	"fmt"
	"os"
)

// 1.文件对象的类型
// 2.获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	FileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("get info failed")
		return
	}
	//fmt.Println(FileInfo)
	fmt.Printf("文件大小是%d8\n", FileInfo.Size())
	fmt.Printf("文件名是%s\n", FileInfo.Name())
}
