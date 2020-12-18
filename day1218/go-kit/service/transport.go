package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req countRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
func EncodeResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
	return json.NewEncoder(w).Encode(&r)
}
