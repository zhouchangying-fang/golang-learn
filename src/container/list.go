package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushFront("zhou")
	l.PushBack("fang")
	f1 := l.Front()
	f2 := f1.Next()
	l.InsertAfter("tao", f2)
	f3 := f1.Next().Next()
	l.InsertAfter("wo", f3)
	fmt.Println(f2.Value)
	fmt.Println(l.Len())
	l.MoveAfter(f1, f2)
	for lf := l.Front(); lf != nil; lf = lf.Next() {
		fmt.Print(lf.Value)
	}

}
