package service

import (
	"errors"

	"github.com/ilhm-rai/go-middleware/entity"
	"github.com/ilhm-rai/go-middleware/helper"
	"github.com/ilhm-rai/go-middleware/model"
	"github.com/ilhm-rai/go-middleware/repository"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) Login(request model.LoginUserRequest) (token string, err error) {
	user, err := service.UserRepository.FindByEmail(request.Email)

	valid := helper.ComparePass([]byte(user.Password), []byte(request.Password))

	if err != nil || !valid {
		err = errors.New("invalid email or password")
		return
	}

	token = helper.GenerateToken(user.ID, user.Email, string(user.Role))

	return
}

func (service *userServiceImpl) FindByEmail(email string) (response model.UserResponse, err error) {
	user, err := service.UserRepository.FindByEmail(email)

	if err != nil {
		return
	}

	response = model.UserResponse{
		Id:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.Role,
	}

	return
}

func (service *userServiceImpl) Register(request model.RegisterUserRequest) (response model.UserResponse, err error) {
	user := entity.User{
		Email:    request.Email,
		FullName: request.FullName,
		Role:     request.Role,
		Password: request.Password,
	}

	newId, err := service.UserRepository.Insert(user)

	if err != nil {
		return
	}

	response = model.UserResponse{
		Id:       newId,
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.Role,
	}

	return
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}
