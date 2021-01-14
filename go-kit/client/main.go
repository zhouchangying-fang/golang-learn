package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
	"golang-learn/go-kit/client/service"
	"io"
	"net/url"
	"os"
	"time"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "192.168.26.10:8500"
	client, err := api.NewClient(config)
	if err != nil {

	}
	client1 := consul.NewClient(client)
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
	}
	tags := []string{"primary"}
	instance := consul.NewInstancer(client1, logger, "service", tags, true)
	f := func(serviceAddr string) (endpoint.Endpoint, io.Closer, error) {
		tag, _ := url.Parse("http://" + serviceAddr)
		return kithttp.NewClient("GET", tag, service.EncodeUserRequest, service.DecodeUserResponse).Endpoint(), nil, nil
	}
	edt := sd.NewEndpointer(instance, f, logger)
	bl := lb.NewRoundRobin(edt)

	for {
		edp, err := bl.Endpoint()
		if err != nil {
			fmt.Println("-----------", err)
		}
		ctx := context.Background()
		res, err := edp(ctx, service.UserRequest{
			Uid: 100,
		})
		if err != nil {
			fmt.Println(err)
		}
		rep := res.(service.UserResponse)
		fmt.Println("------------", rep.UserName)
		time.Sleep(3 * time.Second)
	}

	/*u, _ := url.Parse("http://localhost:8080")
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
	fmt.Println("------------", rep.UserName)*/
}
