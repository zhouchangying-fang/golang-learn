package main

import (
	"container/list"
	"container/ring"
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
	fmt.Println()
	fmt.Println("-----------------ring-------------------")
	r := ring.New(6)
	r1 := ring.New(6)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r1.Value = i
		r1 = r1.Next()
		r = r.Next()
	}
	//r.Link(r1)
	//r = r.Move(5)
	r.Unlink(1)
	r.Do(func(i interface{}) {
		fmt.Print(i, ",")
	})
}
