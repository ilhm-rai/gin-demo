package service

import (
	"github.com/ilhm-rai/go-middleware/entity"
	"github.com/ilhm-rai/go-middleware/model"
	"github.com/ilhm-rai/go-middleware/repository"
)

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func (service *productServiceImpl) DeleteProduct(id uint) (err error) {
	err = service.ProductRepository.DeleteProduct(id)
	return
}

func (service *productServiceImpl) UpdateProduct(id uint, request model.UpdateProductRequest) (err error) {
	product := entity.Product{
		Title:       request.Title,
		Description: request.Description,
		UserID:      request.UserID,
	}
	err = service.ProductRepository.UpdateProduct(id, product)
	return
}

func (service *productServiceImpl) FindProduct(id uint) (response *model.ProductResponse, err error) {
	product, err := service.ProductRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	response = &model.ProductResponse{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		UserID:      product.UserID,
	}

	return response, nil
}

func (service *productServiceImpl) FindProductsByUserId(userId uint) (response []model.ProductResponse, err error) {
	products, err := service.ProductRepository.FindByUserId(userId)

	if err != nil {
		return
	}

	for _, product := range products {
		response = append(response, model.ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			UserID:      product.UserID,
		})
	}

	return
}

func (service *productServiceImpl) FindProducts() (response []model.ProductResponse, err error) {
	products, err := service.ProductRepository.FindAll()

	if err != nil {
		return
	}

	for _, product := range products {
		response = append(response, model.ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			UserID:      product.UserID,
		})
	}

	return
}

func (service *productServiceImpl) CreateProduct(request model.CreateProductRequest) (response model.ProductResponse, err error) {
	product := entity.Product{
		Title:       request.Title,
		Description: request.Description,
		UserID:      request.UserID,
	}

	productId, err := service.ProductRepository.CreateProduct(product)
	if err != nil {
		return
	}

	response = model.ProductResponse{
		ID:          productId,
		Title:       product.Title,
		Description: product.Description,
		UserID:      product.UserID,
	}

	return
}

func NewProductService(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
	}
}
