package main

import (
	"fmt"
	"os"
)

//学生管理系统，实现学生的增、删、查

func main() {
	for {
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("hello")
		case 2:
			fmt.Println("world")
		case 3:
			os.Exit(1)
		default:
			fmt.Println("输入错误")
		}

		// 1.输出提示，让输入相应的选项，进行相应的操作
		// 2.增加学生（学号、姓名）
		// 3.删除学生（根据学号进行删除）
		// 4.查看所有学生的信息
		// 5.给出退出选项

	}
}
