package main

import "fmt"

//标识符：变量名 函数名 类型名 方法名
//Go 语言中如果标识符首字母是大写的，就表示对外部可见

// Dog 这是一个狗的结构体
type Dog struct {
	name string
}

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//方法作用于特定类型的函数
//dog是结构体，接收者d表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s 汪汪汪~\n", d.name)
}

//使用值接收者：传值
func (p person) guonian() {
	p.age++
}

//指针接受者：传内存地址
func (p *person) zhenguonian() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("自由")
}

func main() {
	d1 := newDog("heibei")
	d1.wang()
	p1 := newPerson("zhangsan", 19)
	fmt.Println(p1)
	p1.guonian()
	fmt.Println(p1)
	(&p1).zhenguonian()
	p1.zhenguonian()
	fmt.Println(p1)
	(&p1).dream()
}
