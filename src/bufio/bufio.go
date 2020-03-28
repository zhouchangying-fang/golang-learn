package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	n, t, e := bufio.ScanBytes([]byte("zhoufangtao"), true)
	fmt.Printf("ScanBytes:返回第一个字母---advance:%d,t:%s,err:%v\n", n, t, e)
	n1, t1, e1 := bufio.ScanLines([]byte("zhou\nfangtao"), true)
	fmt.Printf("ScanLines:返回换行前的字符切片---advance:%d,t1:%s,err:%v\n", n1, t1, e1)
	n2, t2, e2 := bufio.ScanRunes([]byte("周芳涛"), true)
	fmt.Printf("ScanRunes:返回第一个字---advance:%d,t1:%s,err:%v\n", n2, t2, e2)
	n3, t3, e3 := bufio.ScanWords([]byte("zhou fang tao"), true)
	fmt.Printf("ScanWords:返回空格前第一个字---n3:%d,t3:%s,err:%v\n", n3, t3, e3)

	ior := strings.NewReader("woaiheini ni")
	r := bufio.NewReader(ior)
	r = bufio.NewReaderSize(ior, 10)
	fmt.Println(r.Buffered())
	n4, e4 := r.Discard(2)
	fmt.Println("r.Discard:---", n4, e4)
	b1, e1 := r.Peek(2)
	fmt.Printf("r.Peek:b1:%s,%v---\n", b1, e1)
	buf := make([]byte, 2)
	n5, e5 := r.Read(buf)
	fmt.Printf("r.Read:---%d,%v,%s\n", n5, e5, buf)
	b2, e2 := r.ReadByte()
	fmt.Println(b2, e2)
	ior2 := strings.NewReader("woaiheini")
	r.Reset(ior2)
	r.UnreadByte()
	r.UnreadRune()
	s, e := r.ReadString('i')
	fmt.Println("r.ReadString:--", s, e)
	fmt.Println("r.Size:--", r.Size())
	r.WriteTo(os.Stdout)

}
