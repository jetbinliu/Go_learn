package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看、新增、删除学生
*/

var (
	allStudent map[int64]*student //声明变量

)

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	// 把所有学生打印出来
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s id:%d\n", k, v.name, v.id)
	}
}

func addStudent() {
	// 想allStudent中添加一个学生
	// 1 .创建一个新学生
	// 1.1获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	// 1.2创造学生(调用构造函数)
	newStu := newStudent(id, name)
	//2.追加到allStudent这个map中
	allStudent[id] = newStu

}

func deleteStudent() {
	// 1.请输入要删除的学生的学号
	var (
		deleteID int64
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&deleteID)
	delete(allStudent, deleteID)
}

func main() {
	allStudent = make(map[int64]*student, 50) //初始化（开辟内存空间）
	for {
		// 1.打印菜单
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出系统
		`)
		fmt.Printf("请输入你要干啥：")
		// 2.等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d 这个选项！\n", choice)
		// 3.执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入错误")
		}

	}
}
