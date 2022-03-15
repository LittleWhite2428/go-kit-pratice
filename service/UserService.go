package service

type IUserService interface {
	GetName(userId int) string
}

type UserService struct {}

func (U UserService ) GetName(userId int) string {
	if userId == 101{
		return "wangyuchao"
	}else {
		return "guest"
	}
}