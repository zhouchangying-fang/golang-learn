package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	UserId int `json:"user_id"`
}
type UserResponse struct {
	UserName string `json:"user_name"`
}

func MakeUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)
		res, _ := svc.GetUserName(req.UserId)
		return UserResponse{
			UserName: res,
		}, nil
	}
}
