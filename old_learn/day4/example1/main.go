package main

import (
	"fmt"
	"time"
	"errors"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {

    // defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 		//alarm.Do()
	// 	}
	// }()

	// b :=0
	// a := 100/b
	// fmt.Println(a)
	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	for {
		test()
		time.Sleep(time.Second)
	}
	//test()
	var a []int
	a = append(a,10,20,30)
	a = append(a,a...)
	fmt.Println(a)


   
	var i int
	fmt.Println(i)

	j := new(int)  //new 返回地址
	*j = 100
	fmt.Println(j)
}