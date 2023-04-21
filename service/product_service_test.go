package service

import (
	"testing"

	"github.com/ilhm-rai/go-middleware/entity"
	"github.com/ilhm-rai/go-middleware/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var productRepository = repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = productServiceImpl{ProductRepository: &productRepository}

func TestFindProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", uint(1)).Return(nil)

	product, err := productService.FindProduct(uint(1))

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
}

func TestFindProductFound(t *testing.T) {
	expected := entity.Product{
		Model: gorm.Model{
			ID: 1,
		},
		Title:       "Test",
		Description: "Lorem ipsum sit",
		UserID:      1,
	}

	productRepository.Mock.On("FindById", uint(1)).Return(expected)

	product, err := productService.FindProduct(uint(1))

	assert.Nil(t, err)
	assert.Equal(t, expected.ID, product.ID)
	assert.Equal(t, expected.Title, product.Title)
}

func TestFindProductsNotFound(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(nil)

	products, err := productService.FindProducts()

	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindProductsFound(t *testing.T) {
	expected := []entity.Product{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Title:       "Test 1",
			Description: "Lorem ipsum",
			UserID:      1,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Title:       "Test 2",
			Description: "Lorem ipsum",
			UserID:      1,
		},
	}
	productRepository.Mock.On("FindAll").Return(expected)

	products, err := productService.FindProducts()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
	assert.Equal(t, expected[0].Title, products[0].Title)
	assert.Equal(t, expected[1].Title, products[1].Title)
}
