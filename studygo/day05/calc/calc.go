package calc

import "fmt"

func init() {
	fmt.Println("import 我时自动调用", pi)
}

const (
	pi = 3.14
)

// Add ... 包中的标识符(变量名、函数名、结构体、接口等)如果首字母是小写的，表示私有(只是当前包中使用)
// Add ... 包中的标识符如果首字母大写，可以被外部调用
func Add(x, y int) int {
	return x + y
}
