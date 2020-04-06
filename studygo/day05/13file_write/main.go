package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_TRUNC|os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed ,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	// write
	fileObj.Write([]byte("蒙蔽了\n"))
	// writeString
	fileObj.WriteString("mengbumeng\n")
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./demo2.txt", os.O_TRUNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("write file failed, err:%v", err)
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	for i := 0; i < 5; i++ {
		//fmt.Print(i)
		wr.WriteString("hello china ")

	}
	wr.Flush()
}

func writeDemo3() {
	str := "hello world\n"
	err := ioutil.WriteFile("./demo3.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed ,err :%v", err)
		return
	}
}

func main() {
	//writeDemo1()
	//writeDemo2()
	writeDemo3()
}
