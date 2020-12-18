package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"golang-learn/day1218/go-kit/service"
	"net/http"
)

func main() {
	srv := service.StringServiceImpl{}
	uppercaseHandler := httptransport.NewServer(
		service.MakeUppercaseEndpoint(srv),
		service.DecodeUppercaseRequest,
		service.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		service.MakeCountEndpoint(srv),
		service.DecodeCountRequest,
		service.EncodeResponse,
	)
	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}
