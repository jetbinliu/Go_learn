package main

import (
	"fmt"
)

func testMap() {
	var a map[string]string  //声明map a
	a = make(map[string]string,10)  //初始化
	a["abc"] = "efg"
	a["abc"] = "a"
	a["d"] = "b"
	fmt.Println(a)
}

func testMap2() {
	a := make(map[string]map[string]string,100)
	a["k1"] = make(map[string]string)
	a["k2"] = make(map[string]string)
	a["k1"]["k1-1"] = "a"
	a["k2"]["k2-1"] = "b"
	fmt.Println(a)
	val, ok := a["k3"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("no")
	}
}

func testMap3() {
	a := make(map[string]string,10)
	a["abc"] = "efg"
	a["abc"] = "a"
	a["d"] = "b"
	delete(a,"d")
	for k,v := range a {
		fmt.Println(k,v)
	}
	
	fmt.Println(len(a))
}


func trans(a map[string]map[string]string) {
	for k,v := range a{
		fmt.Println(k)
		for k1,v1 := range v{
			fmt.Println("\t",k1,v1)
		}
	}
}

func testMap4() {
	a := make(map[string]map[string]string,10)
	a["k1"] = make(map[string]string)
	a["k1"]["k1-1"] = "a"
	a["k1"]["k1-2"] = "b"
	a["k1"]["k1-3"] = "abc"

	a["k2"] = make(map[string]string)
	a["k2"]["k2-1"] = "ka"
	a["k2"]["k2-2"] = "kb"

	trans(a)
	fmt.Println()
	fmt.Println(len(a))
	delete(a,"k2")
	trans(a)
	fmt.Println(len(a))

}

func testMap5() {
	a := make([]map[int]int,10)
	
	a[0] = make(map[int]int)
	a[0][10] = 100
	fmt.Println(a)
}

func testMap6() {

}

func main() {
	//testMap4()
	testMap5()
	
}