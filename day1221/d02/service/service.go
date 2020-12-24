package service

type Service interface {
	Add(a, b int) int
}

type serviceImpl struct {
}

func (st serviceImpl) Add(a, b int) int {
	return a + b
}

type serviceRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}
type serviceResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}
