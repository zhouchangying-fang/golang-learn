package main

import (
	"golang-learn/gokit/stringsvg/service"
	"net/http"
)

func main() {
	upperHandler := service.GetUppercaseHandler()
	http.Handle("/upper", upperHandler)
	http.ListenAndServe(":8080", nil)
}
