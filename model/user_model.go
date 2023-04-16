package model

import "github.com/ilhm-rai/go-middleware/entity"

type RegisterUserRequest struct {
	Email    string      `json:"email" valid:"required~Email is required,email~Invalid email format"`
	FullName string      `json:"full_name" valid:"required~Full name is required"`
	Role     entity.Role `json:"role" valid:"required~Role is required"`
	Password string      `json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type LoginUserRequest struct {
	Email    string `json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password string `json:"password" valid:"required~Password is required"`
}

type UserResponse struct {
	Id       uint        `json:"id"`
	Email    string      `json:"email"`
	FullName string      `json:"full_name"`
	Role     entity.Role `json:"role"`
}
