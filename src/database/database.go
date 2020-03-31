package main

import (
	"context"
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

}
