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
	fmt.Println("递阶创建文件夹", os.MkdirAll("zhou/txt", 777))
	f, e := os.Create("zhou/a.txt")
	fmt.Println("Creat:创建一个O_RDWR（0666）可读写的文件---", f, e)
	fmt.Println("Link:复制文件", os.Link("zhou/a.txt", "zhou/txt/a.txt"))
	s, e := os.Readlink("zhou/txt/a.txt")
	fmt.Println("remove:删除文件", os.Remove("zhou/txt/a.txt"), os.Remove("zhou/a.txt"))
	fmt.Println("Readlink:-s:", s, "e:", e)
	r, w, e := os.Pipe()
	fmt.Println(r, w, e)
	u, e := os.UserCacheDir()
	fmt.Println("UserCacheDir:默认缓存文件夹---", u, e)
	u, e = os.UserConfigDir()
	fmt.Println("UserConfigDir:默认文件配置文件---", u)
	u, e = os.UserHomeDir()
	fmt.Println("UserHomeDir:返回主机用户文件夹---", u, e)

	f, e = os.OpenFile("zhou/b.txt", os.O_RDWR|os.O_CREATE, 0777)
	fmt.Println("OpenFile:打开或创建文件---", f, e)

	n, e := f.Write([]byte("abcd"))
	fmt.Println("Write:---", n, e)

	fr, e := os.Open("zhou/a.txt")
	fmt.Println("Open:打开的文件只能读（ O_RDONLY）--", fr, e)
	defer fr.Close()
	var b = make([]byte, 2)
	n, e = f.Read(b)
	fmt.Println(n, e)
	fmt.Println("Fd:文件描述符---", f.Fd())
	defer f.Close()
	fileInfo, e := f.Readdir(-1)
	fmt.Println("fileInfo:", fileInfo, "err:", e)
	nams, e := f.Readdirnames(-1)
	fmt.Println("nams:", nams, "err:", e)

	rt, e := f.Seek(10, 0)
	fmt.Println("Seek:设置读或写的位置---", rt, e)

	fo, e := f.Stat()
	fmt.Println("Stat:获取文件结构描述---", fo, e)

	fw, _ := f.WriteString("zhoufangtao")
	fmt.Println("Writstring:写入字符串---", fw)

	FileInfo, _ := os.Stat("zhou/a.txt")
	//os.Lstat(也可以读快捷方式)
	fmt.Println("Stat:获取文件信息---", FileInfo)
	fmt.Println(FileInfo.Name(), ".......")

	fmt.Println("IsDir:---", FileInfo.Mode().IsDir())
	fmt.Println("IsRegular:---", FileInfo.Mode().IsRegular())
	fmt.Println("String:权限---", FileInfo.Mode().String())
	//fmt.Println("RemoveAll:",os.RemoveAll("zhou"))
}
