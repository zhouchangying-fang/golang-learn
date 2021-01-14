package service

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

type UserRequest struct {
	UserId int `json:"user_id"`
}
type UserResponse struct {
	UserName string `json:"user_name"`
}

func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, errors.New("too many request")
			}
			return next(ctx, request)
		}
	}
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
