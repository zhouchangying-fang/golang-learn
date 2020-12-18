package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svr StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uppercaseRequest)
		str, err := svr.Uppercase(req.Str)
		if err != nil {
			return uppercaseResponse{"", err}, err
		}
		return uppercaseResponse{Str: str, err: nil}, nil
	}
}
func MakeCountEndpoint(svr StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(countRequest)
		num := svr.Count(req.Str)
		return countResponse{
			Num: num,
		}, nil
	}
}
