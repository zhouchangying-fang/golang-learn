package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uppercaseRequest)
		str, err := svc.Uppercase(req.Str)
		return uppercaseResponse{
			Str: str,
			Err: err,
		}, nil
	}
}
func MakeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(countRequest)
		cnt := svc.Count(req.Str)
		return countResponse{
			Cnt: cnt,
		}, nil
	}
}
