package main

import "fmt"

type student struct {
	id   int64
	name string
}

// 学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudents() {
	// 从s.allStudent这个map中把所有的学生逐个拎出来，stu是具体每一个学生
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d  姓名:%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudents() {
	// 1.根据用户输入的内容创建新的学生
	var (
		stuID   int64
		stuName string
	)
	// 获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	// 根据用户输入创建结构体对象
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	// 2.把新的学生放到s.allStudent这个map中
	fmt.Println("添加成功")
}

// 修改学生
func (s studentMgr) editStudents() {
	// 1.获取用户输入的学号
	var stuID int64
	fmt.Print("请输入要修改的学号")
	fmt.Scanln(&stuID)

	stuObj, ok := s.allStudent[stuID]
	if ok == false {
		fmt.Println("无该学生，请重新输入学号")
		return
	}
	fmt.Printf("你要修改的学生信息如下： 学号：%d , 姓名：%s", stuObj.id, stuObj.name)
	fmt.Print("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	//s.allStudent[stuID].name = newName ??????
	stuObj.name = newName
	s.allStudent[stuID] = stuObj
}

// 删除学生
func (s studentMgr) deleteStudents() {
	// 1.请用户输入要删除学生的Id
	// 2.去map里查找，没有就查无此人，有就删除
	var stuID int64
	fmt.Printf("请输入要删除学生的学号：")
	fmt.Scanln(&stuID)
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Print("查无此人")
		return
	}
	// 3.有就删除
	delete(s.allStudent, stuID)
	fmt.Println("删除成功")
}
