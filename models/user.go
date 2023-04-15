package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/ilhm-rai/go-middleware/helpers"
	"gorm.io/gorm"
)

type Role string

const (
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

type User struct {
	gorm.Model
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email format"`
	Role     Role      `gorm:"type:role_gm;not null"  json:"role" form:"role" valid:"required~Role is required"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

// func (User) TableName() string {
// 	return "user"
// }

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)

	if err != nil {
		return
	}

	u.Password = helpers.HashPass(u.Password)

	return
}
