package main

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
)

func main() {
	fmt.Printf("Drivers:---%s\n", sql.Drivers())
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/golang?charset=utf8")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	err1 := db.Ping()
	if err1 != nil {
		fmt.Println("err1:", err1)
		return
	}
	defer db.Close()
	fmt.Println("Driver:---", db.Driver())
	res, er := db.Exec("INSERT INTO user (name,pwd) VALUES (?,?)", "张三", fmt.Sprintf("%x", md5.Sum([]byte("123456"))))
	fmt.Printf("res:%s,err:%v", res, er)
}
