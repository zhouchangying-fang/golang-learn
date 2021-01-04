package main

import (
	"context"
	"golang-learn/day1221/d01/service"
	"net/http"
)

func main() {
	ctx := context.Background()
	svc := service.ServiceImpl{}
	endpoint := service.MakeServiceImplEndpoint(svc)
	handler := service.MakeHttpHandler(ctx, endpoint)
	http.ListenAndServe(":8080", handler)
}
