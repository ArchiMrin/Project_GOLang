package interfaces

import "github.com/ArchiMrin/Project_GOLang/eCOMM/entities"

type IUser interface {
	Register(user *entities.User) (*entities.RegisterResponse, error)
	//GetRegister(user *entities.User) string
	Login(user *entities.Login) (*entities.LoginResponse, error)
	GetUser(userId string) (*entities.User, error)
	//GetProfile(userId int) *entities.User
	//GetLogout() string
}
