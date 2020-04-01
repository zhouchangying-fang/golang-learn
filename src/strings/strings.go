package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Count:----Count", strings.Count("zhoufangtao", "a"))
	var b strings.Builder
	b.Write([]byte("12345678●■●◆▲△"))
	b.WriteString("abcd")
	b.Grow(2)
	fmt.Println(b.String(), b.Cap())
	r := strings.NewReader("zhoufangtao")
	buf := make([]byte, 4)
	r.Read(buf)
	fmt.Println(r.Len())
	fmt.Printf("%s\n", buf)
	rp := strings.NewReplacer("<", ">", ">", "<")
	fmt.Println(rp.Replace("<zhou>"))
}
