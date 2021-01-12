package main

import (
	"context"
	"fmt"
	kithttp "github.com/go-kit/kit/transport/http"
	"golang-learn/go-kit/client/service"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://localhost:8080")
	client := kithttp.NewClient("GET", u, service.EncodeUserRequest, service.DecodeUserResponse)
	edp := client.Endpoint()
	ctx := context.Background()
	res, err := edp(ctx, service.UserRequest{
		Uid: 100,
	})
	if err != nil {
		fmt.Println(err)
	}
	rep := res.(service.UserResponse)
	fmt.Println("------------", rep.UserName)
}
