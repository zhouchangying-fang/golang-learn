package uitl

import (
	"context"
	"net/http"
)

type MyError struct {
	Code int
	Msg  string
}

func NewMyError(code int, msg string) error {
	return &MyError{
		Code: code,
		Msg:  msg,
	}
}
func (st *MyError) Error() string {
	return st.Msg
}
func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(404)
	w.Write([]byte(err.Error()))
}
