package main

import (
	"context"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"golang-learn/go-kit/log/service"
	"net/http"
	"os"
)

func main() {
	var svc service.Service
	svc = service.ServiceImpl{}
	ctx := context.Background()
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stderr)
		logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}
	svc = service.LoggingMiddleware(logger)(svc)
	edp := service.MakeServiceEndpoint(svc)
	hdl := service.MakeHandler(ctx, edp)
	errChan := make(chan error)
	go func() {
		fmt.Println("start 8080")
		errChan <- http.ListenAndServe(":8080", hdl)
	}()
	fmt.Println(<-errChan)
}
