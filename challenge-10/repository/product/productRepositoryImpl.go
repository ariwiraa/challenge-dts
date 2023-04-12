package repository

import (
	"challange-10/entity"
	"errors"
	"log"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

// Delete implements ProductRepository
func (repository *ProductRepositoryImpl) Delete(product entity.Product) {
	err := repository.DB.Where("id = ?", product.ID).Delete(&product).Error
	if err != nil {
		log.Fatal("error deleting data: ", err.Error())
		return
	}
}

// FindById implements ProductRepository
func (repository *ProductRepositoryImpl) FindById(id uint) (entity.Product, error) {
	var product entity.Product

	err := repository.DB.Preload("User").First(&product, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("product is not found")
		}
		log.Fatal("error getting product data: ", err)
	}
	return product, nil
}

// Update implements ProductRepository
func (repository *ProductRepositoryImpl) Update(product entity.Product, id uint) (entity.Product, error) {
	err := repository.DB.Model(&product).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

// Create implements ProductRepository
func (repository *ProductRepositoryImpl) Create(product entity.Product) (entity.Product, error) {
	err := repository.DB.Debug().Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}
