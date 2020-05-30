package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Go连接MySQL实例

var db *sql.DB // 是一个连接池

func initDB() (err error) {
	// 连接数据库
	dsn := "test:test@@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn) // 千万不要用:=  不校验用户名和密码是否正确
	if err != nil {                  // dsn格式不挣钱
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数s
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 查询单条记录
func queryOne(id int) {
	var u1 user
	// 1.写查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id=?;` // 用？为占位符
	// 2.执行并拿到结果
	// 从连接池里面拿一个连接出来去数据库查询单条记录，1表示id=1 ，必须对rowObj进行scan，scan自带释放连接
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	fmt.Printf("u1:%#v\n", u1)
}

func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init DB failed, err:", err)
		return
	}
	fmt.Println("2连接数据库成功！")
	queryOne(1)
}
