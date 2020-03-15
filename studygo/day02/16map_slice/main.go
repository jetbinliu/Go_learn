package main

import "fmt"

//切片中的元素为map
func mapInSlice() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	//对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "zhangsan"
	mapSlice[0]["password"] = "passwd"
	mapSlice[0]["address"] = "china"
	for index, value := range mapSlice {

		if value == nil {
			continue
		}
		fmt.Printf("index:%d value:%v\n", index, value)
		for k, v := range value {
			fmt.Println("\t", k, v)
		}

	}
}

//map中的值为slice
func sliceInMap() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "china"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2) //没有该key，则先进行初始化
	}
	value = append(value, "北京", "shanghai")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func main() {
	//mapInSlice()
	sliceInMap()
}
