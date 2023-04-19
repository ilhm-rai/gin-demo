package repository

import (
	"github.com/ilhm-rai/go-middleware/entity"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *productRepositoryImpl) DeleteProduct(id uint) (err error) {
	err = repository.DB.Delete(&entity.Product{}, id).Error
	return
}

func (repository *productRepositoryImpl) UpdateProduct(id uint, product entity.Product) (err error) {
	err = repository.DB.Model(entity.Product{}).Where("id = ?", id).Updates(product).Error
	return
}

func (repository *productRepositoryImpl) FindAll() (products []entity.Product, err error) {
	err = repository.DB.Find(&products).Error
	return
}

func (repository *productRepositoryImpl) FindById(id uint) (product entity.Product, err error) {
	err = repository.DB.Where("id = ?", id).First(&product).Error
	return
}

func (repository *productRepositoryImpl) FindByUserId(userId uint) (products []entity.Product, err error) {
	err = repository.DB.Where("user_id = ?", userId).Find(&products).Error
	return
}

func (repository *productRepositoryImpl) CreateProduct(product entity.Product) (id uint, err error) {
	err = repository.DB.Create(&product).Error
	return product.ID, err
}

func NewProductRepository(database *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		DB: database,
	}
}
