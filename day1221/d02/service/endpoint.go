package service

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	ktp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func MakeServiceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		requestError := errors.New("request error")
		req := request.(serviceRequest)
		rst := 0
		if strings.EqualFold(req.RequestType, "Add") {
			rst = svc.Add(req.A, req.B)
		} else {
			return "", requestError
		}
		return serviceResponse{
			Result: rst,
			Error:  nil,
		}, nil

	}
}

func MakeHandler(ctx context.Context, edp endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(
		ktp.NewServer(
			edp,
			decodeServiceRequest,
			encodeServiceResponse,
		),
	)
	return r
}
