package service

type UserService interface {
	GetUserName(userId int) (string, error)
}
type UserServiceImpl struct {
}

func (st UserServiceImpl) GetUserName(userId int) (string, error) {
	if userId == 100 {
		return "heini", nil
	}
	return "not know", nil
}
