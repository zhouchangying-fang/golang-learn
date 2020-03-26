package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("返回当前计算机所有环境变量:", os.Environ())
	ex, err := os.Executable()
	fmt.Println("返回当前进程入口文件地址：", ex, err)
	fmt.Println(os.IsPathSeparator('/'))
	os.Link("e", "w")
	os.Setenv("name", "zhoufangtao")
	l, bool := os.LookupEnv("name")
	fmt.Println(l, bool)
	fmt.Println("在当前gopath下新建zhou文件夹", os.Mkdir("zhou", 777))
	fmt.Println("递阶创建文件夹", os.MkdirAll("E:/zhou/golang/zhou/txt", 777))
	fmt.Println("Link:复制文件", os.Link("zhou/a.txt", "zhou/txt/a.txt"))
	s, e := os.Readlink("zhou/txt/a.txt")
	fmt.Println("remove:删除文件", os.Remove("zhou/txt/a.txt"))
	fmt.Println("Readlink:-s:", s, "e:", e)
	r, w, e := os.Pipe()
	fmt.Println(r, w, e)
	//fmt.Println("RemoveAll:",os.RemoveAll("zhou"))
}
