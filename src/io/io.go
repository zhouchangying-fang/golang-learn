package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("Copy:--\n")
	io.Copy(os.Stdout, r1)

	r2 := strings.NewReader("NewReader:---\n")
	b := make([]byte, 1)
	io.CopyBuffer(os.Stdout, r2, b)

	r3 := strings.NewReader("CopyN:复制N个长度---\n")
	io.CopyN(os.Stdout, r3, 2)

	r, w := io.Pipe()
	fmt.Println("Pipe:---", r, w)

	r4 := strings.NewReader("ReadAtLeast\n")
	b4 := make([]byte, 6)
	n, e := io.ReadAtLeast(r4, b4, 1)
	fmt.Printf("ReadAtLeast:---%v,%v,%s\n", n, e, b4)

	r5 := strings.NewReader("ReadFull ")
	buf5 := make([]byte, 3)
	n, e = io.ReadFull(r5, buf5)
	fmt.Printf("ReadFull:---%v,%v,%s\n", n, e, buf5)

	io.WriteString(os.Stdout, "WriteString\n")

	fmt.Println("--------io-ioutil-----------------io-ioutil--------------------")

	r6 := strings.NewReader("NopCloser")
	ioutil.NopCloser(r6)
	buf, e := ioutil.ReadAll(r6)
	fmt.Printf("ReadAll:读取所有---%s,err:%v\n", buf, e)
	F, _ := ioutil.ReadDir("zhou")
	for _, f := range F {
		fmt.Println("ReadDir:读取文件下列表---", f.Name())
	}
	F2, e2 := ioutil.ReadFile("zhou/a.txt")
	fmt.Printf("ReadFile:---%s,err:%v\n", F2, e2)
	ioutil.TempDir("zhou", "zhoufangtao")

}
