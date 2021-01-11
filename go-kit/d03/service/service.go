package service

import "errors"

type Service interface {
	Add(a, b int) int
	Subtract(a, b int) int
	Multiply(a, b int) int
	Divide(a, b int) (int, error)
}

type ServiceImpl struct {
}

func (st ServiceImpl) Add(a, b int) int {
	return a + b
}
func (st ServiceImpl) Subtract(a, b int) int {
	return a - b
}
func (st ServiceImpl) Multiply(a, b int) int {
	return a * b
}
func (st ServiceImpl) Divide(a, b int) (int, error) {
	if a == 0 {
		return 0, errors.New("the divide can not be zero!")
	}
	return a / b, nil
}
