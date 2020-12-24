package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"golang-learn/day1221/d02/service"
	"net/http"
)

func main() {
	var svc service.Service
	ctx := context.Background()
	var log log.Logger
	svc = service.LoggingMiddleware(log)(svc)
	edp := service.MakeServiceEndpoint(svc)
	hdl := service.MakeHandler(ctx, edp)
	errChan := make(chan error)
	go func() {
		fmt.Println("start 8080")
		errChan <- http.ListenAndServe(":8080", hdl)
	}()
	fmt.Println(<-errChan)
}
