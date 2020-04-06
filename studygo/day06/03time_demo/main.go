package main

import (
	"fmt"
	"time"
)

// 时间
func f1() {
	startTime := time.Now().UnixNano()
	startTime2 := time.Now()
	fmt.Println(startTime)
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(now.UnixNano() / 1000000)

	// time.Unix()
	ret := time.Unix(1586142599, 0)
	fmt.Println(ret)
	ret1, _ := time.Parse("2006/01/02 15 04 05", "2020/04/06 11 28 18")
	fmt.Println(ret1)
	fmt.Println("ret1:", ret1.Unix())

	//时间间隔
	fmt.Println(time.Millisecond + time.Second)
	// now + 1小时
	fmt.Println(now.Add(time.Hour))
	// 定时器
	// timer := time.Tick(time.Second * 2)
	// for t := range timer {
	// 	fmt.Println(t)
	// }
	// 格式化时间
	fmt.Println(now.Format("2006/01/02 15 04 05"))
	fmt.Println(now.Format("2006/01/02 15 04 05.0000"))
	fmt.Println(now.Format("2006/01/02 15 04 05.999"))
	fmt.Println(now.Format("2006/01/02 03 04 05 PM"))
	n := 10
	time.Sleep(time.Duration(n) * time.Second)
	endTime := time.Now().UnixNano()
	fmt.Println(endTime)
	execMicroTime := (endTime - startTime) / 1000
	fmt.Println(execMicroTime)
	endTime2 := time.Now()
	now2 := endTime2.Sub(startTime2)

	fmt.Println(now2)
}

// 时区
func f2() {
	now := time.Now() //获取本地的时间
	fmt.Println(now)
	// 明天的这个时间
	// 按照指定时区的格式解析一个时间
	t1, _ := time.Parse("2006-01-02 15:04:05", "2020-04-06 12:13:38")
	fmt.Println(t1)
	// 按照东八区时区的格式解析一个时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load loc failed", err)
	}
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-04-07 12:23:38", loc)
	if err != nil {
		fmt.Println("load loc failed", err)
	}

	fmt.Println(timeObj)

	td := timeObj.Sub(now)
	fmt.Println(td)
}
func main() {
	f2()
}
