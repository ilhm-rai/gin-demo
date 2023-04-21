package service

import (
	"testing"

	"github.com/ilhm-rai/go-middleware/entity"
	"github.com/ilhm-rai/go-middleware/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ProductServiceTestSuite struct {
	suite.Suite
	ProductRepository repository.ProductRepositoryMock
	ProductService    ProductService
}

func (suite *ProductServiceTestSuite) SetupTest() {
	suite.ProductRepository = repository.ProductRepositoryMock{Mock: mock.Mock{}}
	suite.ProductService = &productServiceImpl{ProductRepository: &suite.ProductRepository}
}

func (suite *ProductServiceTestSuite) TestFindProductNotFound() {
	suite.ProductRepository.Mock.On("FindById", uint(1)).Return(nil)

	product, err := suite.ProductService.FindProduct(uint(1))

	assert.Nil(suite.T(), product)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "record not found", err.Error())
}

func (suite *ProductServiceTestSuite) TestFindProductFound() {
	expected := entity.Product{
		Model: gorm.Model{
			ID: 1,
		},
		Title:       "Test",
		Description: "Lorem ipsum sit",
		UserID:      1,
	}

	suite.ProductRepository.Mock.On("FindById", uint(1)).Return(expected)

	product, err := suite.ProductService.FindProduct(uint(1))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.ID, product.ID)
	assert.Equal(suite.T(), expected.Title, product.Title)
}

func (suite *ProductServiceTestSuite) TestFindProductsNotFound() {
	suite.ProductRepository.Mock.On("FindAll").Return(nil)

	products, err := suite.ProductService.FindProducts()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, len(products))
}

func (suite *ProductServiceTestSuite) TestFindProductsFound() {
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
	suite.ProductRepository.Mock.On("FindAll").Return(expected)

	products, err := suite.ProductService.FindProducts()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(products))
	assert.Equal(suite.T(), expected[0].Title, products[0].Title)
	assert.Equal(suite.T(), expected[1].Title, products[1].Title)
}

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}
