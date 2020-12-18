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

func (st StringServiceImpl) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("string is empty")
	}
	return strings.ToUpper(s), nil
}
func (st StringServiceImpl) Count(s string) int {
	return len(s)
}

type uppercaseRequest struct {
	Str string `json:"str"`
}
type uppercaseResponse struct {
	Str string `json:"str"`
	err error  `json:"err,omitempty"`
}

type countRequest struct {
	Str string `json:"str"`
}
type countResponse struct {
	Num int `json:"num"`
}
