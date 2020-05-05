package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串中解析对应的数字
	str := "10000"
	strInt, _ := strconv.Atoi(str)
	strInt2, _ := strconv.ParseInt(str, 10, 32)
	fmt.Printf("%#v %#v\n", strInt, strInt2)

	// 把数字转换成字符串类型
	i := int32(97)
	ret2 := fmt.Sprintf("%d", i)
	ret3 := strconv.Itoa(int(i))
	fmt.Printf("%#v\n", ret2)
	fmt.Printf("%#v\n", ret3)

	// 从字符串中解析出布尔值
	boolStr := "10"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)
}
