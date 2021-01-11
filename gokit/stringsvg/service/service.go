package service

import (
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringServiceImpl struct {
}

func (st StringServiceImpl) Uppercase(str string) (string, error) {
	if str == "" {
		return "", errors.New("str is empty")
	}
	return strings.ToUpper(str), nil
}
func (st StringServiceImpl) Count(str string) int {
	return len(str)
}

type uppercaseRequest struct {
	Str string `json:"str"`
}
type uppercaseResponse struct {
	Str string `json:"str"`
	Err error  `json:"err,omitempty"`
}
type countRequest struct {
	Str string `json:"str"`
}
type countResponse struct {
	Cnt int `json:"cnt"`
}
