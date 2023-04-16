package entity

import (
	"github.com/ilhm-rai/go-middleware/helper"
	"gorm.io/gorm"
)

type Role string

const (
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

type User struct {
	gorm.Model
	FullName string    `gorm:"not null"`
	Email    string    `gorm:"not null;uniqueIndex"`
	Role     Role      `gorm:"type:role_gm;not null"`
	Password string    `gorm:"not null"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (u *User) BeforeCreate(tx *gorm.DB) {
	u.Password = helper.HashPass(u.Password)
}
