package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

func main() {
	fmt.Println()
	fmt.Println("---------------md5------------")
	h := md5.New()
	io.WriteString(h, "abc")
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Printf("%x\n", md5.Sum([]byte("abc")))
	n, e := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(n, e)
	p, e := rand.Prime(rand.Reader, 4)
	fmt.Println(p, e)
	n2, e2 := rand.Read(make([]byte, 10))
	fmt.Println(n2, e2)
}
