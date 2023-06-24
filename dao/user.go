package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	g "zhihu/global"
)

var err error

// AddUser 向数据库中添加用户
func AddUser(username, password string) error {
	// 准备插入语句
	insertQuery := "INSERT INTO users (username, password) VALUES (?, ?)"

	// 执行插入操作
	_, err = g.MysqlDB.Exec(insertQuery, username, password)
	if err != nil {
		fmt.Printf("Add failed, err:%v\n", err)
		return err
	}
	log.Println("Add success")
	return nil
}

// SelectUser 检查用户是否存在于数据库中
func SelectUser(username string) (bool, error) {
	// 准备查询语句
	selectQuery := "SELECT COUNT(*) FROM users WHERE username = ?"
	var count int

	// 执行查询并将结果存储到 count 变量中
	err = g.MysqlDB.QueryRow(selectQuery, username).Scan(&count)
	if err != nil {
		return false, err
	}

	// 如果 count 大于 0，则用户存在
	// 若没有这个用户返回 false，反之返回 true
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// SelectPasswordFromUsername 根据用户名从数据库中获取密码
func SelectPasswordFromUsername(username string) (string, error) {
	// 准备查询语句
	selectQuery := "SELECT password FROM users WHERE username = ?"
	var password string

	// 执行查询并将结果存储到 password 变量中
	err = g.MysqlDB.QueryRow(selectQuery, username).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}
