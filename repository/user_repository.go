package repository

import "github.com/ilhm-rai/go-middleware/entity"

type UserRepository interface {
	Insert(user entity.User) (uint, error)

	FindByEmail(email string) (entity.User, error)
}
