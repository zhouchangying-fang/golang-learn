package service

import (
	"golang-learn/go-kit/ixigua/uitl"
	"strconv"
)

type UserService interface {
	GetUserName(userId int) (string, error)
}
type UserServiceImpl struct {
}

func (st UserServiceImpl) GetUserName(userId int) (string, error) {
	if userId == 100 {
		return "heini" + strconv.Itoa(uitl.Port), nil
	}
	return "not know", nil
}
