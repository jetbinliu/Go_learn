package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// ini配置文件解析器

// MysqlConfig MySQL结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `int:"port"`
	Username string `ini:"username"`
	Password string `init:"password"`
}

// RedisConfig Redis结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `init:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// Config ...
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 做参数的校验
	// 0.1 传进来的data参数必须是指针类型（因为需要再函数中 对其赋值）
	t := reflect.TypeOf(data)
	//fmt.Println(t, t.Elem().Kind(), t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer") // 格式化输出后返回一个error类型
		//err = fmt.Errorf("data should be a pointer") // 格式化输出后返回一个error类型
		return
	}
	// 0.2 传进来的data参数必须是结构体类型指针
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct")
		return
	}
	// 1. 读文件得到字节类型的数据
	fileObj, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(fileObj) // 将字节类型文件内容转换成字符串
	lineSlice := strings.Split(string(fileObj), "\n")
	fmt.Printf("%#v\n", lineSlice)
	// 2. 一行一行的读数据
	for idx, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)

		// 2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[]开头的表示是节（section）
		if strings.HasPrefix(line, "[") || strings.HasPrefix(line, "]") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 包这一行首尾的[]去掉，取中间的内容并把首尾空格去掉 拿到内容
			if len(strings.TrimSpace(line[1:len(line)-1])) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体

		} else {
			// 2.3 如果不是[]开头，就是=分割的键值对
		}
	}

	return
}

func main() {
	var cfg MysqlConfig
	//test := 10
	err := loadIni("./config.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini file failed,err:%v\n", err)
		return
	}
	fmt.Println(cfg)
	//fmt.Println(cfg.Address, cfg.Port, cfg.Username, cfg.Password)
}
