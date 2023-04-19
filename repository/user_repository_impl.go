package repository

import (
	"github.com/ilhm-rai/go-middleware/entity"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *userRepositoryImpl) FindByEmail(email string) (user entity.User, err error) {
	err = repository.DB.Where("email = ?", email).First(&user).Error
	return
}

func (repository *userRepositoryImpl) Insert(user entity.User) (id uint, err error) {

	err = repository.DB.Create(&user).Error
	return user.ID, err
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: database,
	}
}
