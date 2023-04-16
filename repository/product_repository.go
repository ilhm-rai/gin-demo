package repository

import (
	"github.com/ilhm-rai/go-middleware/entity"
)

type ProductRepository interface {
	CreateProduct(product entity.Product) (id uint, err error)
	UpdateProduct(id uint, product entity.Product) (err error)
	DeleteProduct(id uint) (err error)
	FindAll() (products []entity.Product, err error)
	FindById(id uint) (product entity.Product, err error)
	FindByUserId(userId uint) (products []entity.Product, err error)
}
