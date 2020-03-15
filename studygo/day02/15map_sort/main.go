package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {

		value := rand.Intn(100)
		key := fmt.Sprintf("stu%03d", rand.Intn(1000))

		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	//排序
	//1.取出map中的所有key放入切片
	var keys = make([]string, 0, 200) //初始化一个切片
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//2.对切片进行排序
	sort.Strings(keys)
	//fmt.Println(keys)
	//3.按照排序后的key遍历map
	for _, v := range keys {
		fmt.Println(v, scoreMap[v])
	}
}
