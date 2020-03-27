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
	fmt.Println("Mkdir:在当前gopath下新建zhou文件夹", os.Mkdir("zhou", 777))
	fmt.Println("MkdirAll:递阶创建文件夹", os.MkdirAll("zhou/txt", 777))
	fmt.Println("Link:复制文件", os.Link("zhou/a.txt", "zhou/txt/a.txt"))
	s, e := os.Readlink("zhou/txt/a.txt")
	fmt.Println("Readlink:-s:", s, "e:", e)
	fmt.Println("Rename:移动文件---", os.Rename("zhou/a.txt", "zhou/txt/a.txt"))
	fmt.Println("Symlink:创建快捷方式", os.Symlink("zhou/txt/a.txt", "zhou/a.txt"))
	fmt.Println("remove:删除文件", os.Remove("zhou/a.txt"))
	r, w, e := os.Pipe()
	fmt.Println(r, w, e)
	fmt.Println("TempDir:默认临时文件夹", os.TempDir())
	fmt.Println("Truncata:", os.Truncate("zhou/txt/a.txt", 1024))
	fmt.Println("Setenv:设置env变量---", os.Setenv("name", "zhoufangtao"))
	fmt.Println("Unsetenv:删除env变量---", os.Unsetenv("name"))
	//fmt.Println("RemoveAll:",os.RemoveAll("zhou"))
}
