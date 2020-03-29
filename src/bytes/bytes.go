package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Compare:---", bytes.Compare([]byte("abc"), []byte("abd")))
	fmt.Println("Contains:---", bytes.Contains([]byte("abc"), []byte("a")))
	fmt.Println("ContainsAny:--", bytes.ContainsAny([]byte("abc"), "abd"))
	fmt.Println("ContainsRune:（Rune是一个字节）---", bytes.ContainsRune([]byte("abc"), 'c'))
	fmt.Println("Count:---", bytes.Count([]byte("abcda"), []byte("a")))
	fmt.Println("Equal:---", bytes.Equal([]byte("abc"), []byte("abc")))
	fmt.Println("EqualFold:忽略大小写---", bytes.EqualFold([]byte("HEINI"), []byte("heini")))
	fmt.Printf("Fields:删除空格---%q\n", bytes.Fields([]byte(" wo ai hei ni ")))
	fmt.Printf("FieldsFunc:为真时删除---%s\n", bytes.FieldsFunc([]byte(" wo,ai!hei@ni"), func(r rune) bool {
		return !unicode.IsLetter(r)
	}))
	fmt.Println("HasPrefix:---", bytes.HasPrefix([]byte("woaiheini"), []byte("wo")))
	fmt.Println("HasSuffix:---", bytes.HasSuffix([]byte("woaiheini"), []byte("wo")))
	fmt.Println("Index:---", bytes.Index([]byte("woaiheini"), []byte("ai")))
	fmt.Println("IndexAny:---", bytes.IndexAny([]byte("woaiheini"), "lllai"))
	fmt.Println("IndexByte:---", bytes.IndexByte([]byte("woaiheini"), 'i'))
	fmt.Println("IndexRune:---", bytes.IndexRune([]byte("woaiheini"), 'i'))
	fmt.Printf("Join:---%s\n", bytes.Join([][]byte{[]byte("wo"), []byte("ai"), []byte("ni")}, []byte(",")))
	fmt.Println("LastIndex:---", bytes.LastIndex([]byte("woaiheini"), []byte("i")))
	fmt.Printf("Repeat:---%s\n", bytes.Repeat([]byte("ai"), 3))
	fmt.Printf("Replace:---%s\n", bytes.Replace([]byte("woaiheiheiheini"), []byte("hei"), []byte(""), 2))
	fmt.Printf("Split:---%s\n", bytes.Split([]byte("wo,ai,ni"), []byte(",")))
	fmt.Printf("SplitAft:带分割符的---%s\n", bytes.SplitAfter([]byte("wo,ai,ni"), []byte(",")))
	fmt.Printf("Title:---%s\n", bytes.Title([]byte("wo ai hei ni")))
	fmt.Printf("ToLower:---%s\n", bytes.ToLower([]byte("WOAIheini")))

}
