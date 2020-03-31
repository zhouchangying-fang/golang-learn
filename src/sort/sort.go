package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []float64{1.4, 6.8, 2.1, 64.58, 0.13}
	sort.Float64s(s)
	fmt.Println("Float64s:---", s)
	fmt.Println("Float64sAreSorted:---", sort.Float64sAreSorted(s))
	a := []int{1, 2, 7, 8, 6, 3, 4, 5}
	i := sort.Search(len(a), func(i int) bool {
		return a[i] >= 5
	})
	fmt.Println(a[i])
	if i < len(a) && a[i] == 5 {
		fmt.Println(a[i])
	} else {
		fmt.Println("not found")
	}
}
