package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	b = strconv.AppendQuoteRuneToGraphic(b, 'д')
	fmt.Printf("%s\n", b)
}
