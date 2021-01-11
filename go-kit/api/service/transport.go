package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func decodeServiceImpl(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, errors.New("ErrorBadRequest")
	}
	ra, ok := vars["a"]
	if !ok {
		return nil, errors.New("ErrorBadRequest")
	}
	rb, ok := vars["b"]
	if !ok {
		return nil, errors.New("ErrorBadRequest")
	}
	a, _ := strconv.Atoi(ra)
	b, _ := strconv.Atoi(rb)
	return ServiceImplRequest{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil
}
func encodeServiceImpl(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(
		khttp.NewServer(
			endpoint,
			decodeServiceImpl,
			encodeServiceImpl,
		),
	)
	return r
}
