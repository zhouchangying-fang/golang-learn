package service

type UserRequest struct {
	Uid int `json:"uid"`
}
type UserResponse struct {
	UserName string `json:"user_name"`
}
