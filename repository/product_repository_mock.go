package repository

import (
	"github.com/ilhm-rai/go-middleware/entity"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) DeleteProduct(id uint) (err error) {
	return
}

func (repository *ProductRepositoryMock) UpdateProduct(id uint, product entity.Product) (err error) {
	return
}

func (repository *ProductRepositoryMock) FindAll() (products []entity.Product, err error) {
	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return []entity.Product{}, nil
	}

	products = arguments.Get(0).([]entity.Product)

	return products, nil
}

func (repository *ProductRepositoryMock) FindByUserId(userId uint) (products []entity.Product, err error) {
	return
}

func (repository *ProductRepositoryMock) CreateProduct(product entity.Product) (id uint, err error) {
	return
}

func (repository *ProductRepositoryMock) FindById(id uint) (product entity.Product, err error) {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return product, gorm.ErrRecordNotFound
	}

	product = arguments.Get(0).(entity.Product)

	return product, nil
}
