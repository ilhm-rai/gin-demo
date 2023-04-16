package service

import "github.com/ilhm-rai/go-middleware/model"

type UserService interface {
	Register(request model.RegisterUserRequest) (response model.UserResponse, err error)
	FindByEmail(email string) (response model.UserResponse, err error)
	Login(request model.LoginUserRequest) (token string, err error)
}
