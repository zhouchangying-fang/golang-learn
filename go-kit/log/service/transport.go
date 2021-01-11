package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func decodeServiceRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestError := errors.New("request error")
	a, ok := vars["a"]
	if !ok {
		return "", requestError
	}
	ai, _ := strconv.Atoi(a)
	b, ok := vars["b"]
	if !ok {
		return "", requestError
	}
	bi, _ := strconv.Atoi(b)
	t, ok := vars["type"]
	if !ok {
		return "", requestError
	}
	return serviceRequest{
		t,
		ai,
		bi,
	}, nil
}

func encodeServiceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
