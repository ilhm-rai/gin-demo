package service

import "github.com/ilhm-rai/go-middleware/model"

type ProductService interface {
	CreateProduct(request model.CreateProductRequest) (response model.ProductResponse, err error)
	UpdateProduct(id uint, request model.UpdateProductRequest) (err error)
	DeleteProduct(id uint) (err error)
	FindProduct(productId uint) (product model.ProductResponse, err error)
	FindProducts() (products []model.ProductResponse, err error)
	FindProductsByUserId(userId uint) (products []model.ProductResponse, err error)
}
