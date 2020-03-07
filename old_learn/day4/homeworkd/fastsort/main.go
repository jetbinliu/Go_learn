package main

import (
	"fmt"
)

func fsort(a []int,left int,right int) {
	if left >= right {
		return
	}
	val := a[left]	//确定vale所在的位置
	k := left
	for i := left+1; i <= right;i++{
		if (a[i] < val){
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}

	a[k] = val
	fsort(a,left,k-1)
	fsort(a,k+1,right)
}


func main() {
	b := [...]int{4,2,1,5,17,6}
	fsort(b[:],0 ,len(b)-1)
	fmt.Println(b)
}