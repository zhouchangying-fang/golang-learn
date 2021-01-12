package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func EncodeUserRequest(_ context.Context, r *http.Request, re interface{}) error {
	req := re.(UserRequest)
	r.URL.Path += "/user/" + strconv.Itoa(req.Uid)
	return nil
}
func DecodeUserResponse(_ context.Context, w *http.Response) (response interface{}, err error) {
	var rep UserResponse
	if err := json.NewDecoder(w.Body).Decode(&rep); err == nil {
		return rep, nil
	} else {
		return nil, err
	}

}
