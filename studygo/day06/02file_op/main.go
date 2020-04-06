package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作
func f1() {
	fileObj, err := os.Open("./main.go")
	fmt.Println(fileObj)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer fileObj.Close()
	file, err := ioutil.ReadFile("./main.go")
	fmt.Println(string(file))
}

func f2() {
	// 打开要操纵的文件
	fileObj, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open filed err,err:", err)
		return
	}
	//defer fileObj.Close()

	// 因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./tmp.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("create tmp file failed,err:", err)
		return
	}
	//defer tmpFile.Close()

	// 读原文件，写入临时文件

	//fileObj.Seek(2, 0) // 光标移动

	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	// 写临时文件
	tmpFile.Write(ret[:n])
	// 再写入要插入的内容
	tmpFile.Write([]byte{'c'})
	// 接着把原文件后续内容写到临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Println("read failed,err", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	// 源文件后续的也写入了临时文件中
	// 删除原文件
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./tmp.txt", "./test.txt")

	//fmt.Println(string(ret[:n]))
}

func main() {
	//f1()
	f2()
}
