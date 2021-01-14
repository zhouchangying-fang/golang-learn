package uitl

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
