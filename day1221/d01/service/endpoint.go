package service

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func MakeArithmeticEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)
		var (
			res, a, b int
			calError  error
		)
		a = req.A
		b = req.B
		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Substract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, calError = svc.Divide(a, b)
		} else {
			return nil, errors.New("不知道的请求类型")
		}
		return ArithmeticResponse{
			res,
			calError,
		}, nil
	}
}

func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(
		kithttp.NewServer(
			endpoint,
			decodeArithmeticRequest,
			encodeArithmeticResponse,
		),
	)
	return r
}
