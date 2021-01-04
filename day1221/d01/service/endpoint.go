package service

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"strings"
)

type ServiceImplRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}
type ServiceImplResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

func MakeServiceImplEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ServiceImplRequest)
		var (
			res  int
			a    = req.A
			b    = req.B
			errs error
		)
		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Subtract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, errs = svc.Divide(a, b)
		} else {
			return nil, errors.New("not know request type")
		}
		return ServiceImplResponse{
			Result: res,
			Error:  errs,
		}, nil
	}
}
