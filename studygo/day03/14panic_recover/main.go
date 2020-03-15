package main

import "fmt"

//panic 和recover

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放连接")
	}()
	panic("出现了严重的错误") //让程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	fmt.Printf("%04d\n", 100)
	funcA()
	funcB()
	funcC()

}
