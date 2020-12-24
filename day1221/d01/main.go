package main

import (
	"context"
	"fmt"
	"golang-learn/day1221/d01/service"
	"net/http"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)
	svc := service.ArithmeticService{}
	endpoint := service.MakeArithmeticEndpoint(svc)
	r := service.MakeHttpHandler(ctx, endpoint)
	go func() {
		fmt.Println("Http Server start at port:8080")
		errChan <- http.ListenAndServe(":8080", r)
	}()
	fmt.Println(<-errChan)
}
