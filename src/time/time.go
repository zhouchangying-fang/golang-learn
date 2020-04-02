package main

import (
	"fmt"
	"time"
)

var c chan int

func main() {
	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(2 * time.Second):
		fmt.Println("time out")
	}
	for n := range time.Tick(time.Second) {
		fmt.Println(n)
		break
	}
	hours, _ := time.ParseDuration("10h")
	fmt.Println(hours)
	fh, _ := time.ParseDuration("4h30m")
	fmt.Println(fh.Hours())
	fmt.Println(time.Second.Microseconds())
}
