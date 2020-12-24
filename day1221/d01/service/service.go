package service

import "errors"

type Service interface {
	Add(a, b int) int

	Subtract(a, b int) int

	Multiply(a, b int) int

	Divide(a, b int) (int, error)
}

type ArithmeticService struct{}

func (st ArithmeticService) Add(a, b int) int {
	return a + b
}
func (st ArithmeticService) Subtract(a, b int) int {
	return a - b
}
func (st ArithmeticService) Multiply(a, b int) int {
	return a * b
}
func (st ArithmeticService) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b is can not be zero")
	}
	return a / b, nil
}

type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}
type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}
