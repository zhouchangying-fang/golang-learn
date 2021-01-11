package main

import (
	"context"
	"golang-learn/day1221/gokit-api/service"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

//go-kit 限流器
func main() {
	ctx := context.Background()
	svc := service.ServiceImpl{}
	rat := rate.NewLimiter(rate.Every(1*time.Second), 3)
	edp1 := service.MakeServiceImplEndpoint(svc)
	edp := service.NewTokenBucketLimitterWithBuildIn(rat)(edp1)
	handler := service.MakeHttpHandler(ctx, edp)
	http.ListenAndServe(":8080", handler)
}
