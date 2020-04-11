package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller() failed,err:")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file) // 06 runtime_demo/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line) // 9
}

func main() {
	f1()
}
