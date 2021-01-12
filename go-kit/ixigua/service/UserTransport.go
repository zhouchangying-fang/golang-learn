package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserServiceRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		id, _ := strconv.Atoi(uid)
		return UserRequest{
			UserId: id,
		}, nil
	}
	return nil, errors.New("not find")
}

func EncodeUserServiceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
