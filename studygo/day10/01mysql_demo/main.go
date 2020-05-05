package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Go连接MySQL实例

func main() {
	// 连接数据库
	dsn := "test:test@@tcp(127.0.0.1:3306)/goday10"
	db, err := sql.Open("mysql", dsn) // 不校验用户名和密码是否正确
	if err != nil {                   // dsn格式不挣钱
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功！")
}
