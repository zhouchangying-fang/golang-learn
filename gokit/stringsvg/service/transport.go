package service

import (
	"context"
	"encoding/json"
	"fmt"
	kithttp "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"net/http"
)

func decodeUppercaseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	b, e := ioutil.ReadAll(r.Body)
	fmt.Println(string(b), e)
	fmt.Println(r.Method)
	var req uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func GetUppercaseHandler() *kithttp.Server {
	return kithttp.NewServer(
		MakeUppercaseEndpoint(StringServiceImpl{}),
		decodeUppercaseRequest,
		encodeResponse,
	)
}
