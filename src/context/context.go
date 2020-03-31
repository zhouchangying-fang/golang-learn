package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	fmt.Println("---------WithDeadline---------")
	d := time.Now().Add(1 * time.Second)
	ctx1, cancel1 := context.WithDeadline(context.Background(), d)
	defer cancel1()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("abc")
	case <-ctx1.Done():
		fmt.Println(ctx1.Err())
	}
	fmt.Println("---------WithValue---------")
	type fkey string
	f := func(ctx context.Context, k fkey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("not found key:", k)
	}
	k := fkey("language")
	ctx4 := context.WithValue(context.Background(), k, "Go")
	f(ctx4, k)

}
