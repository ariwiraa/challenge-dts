package service

import "challange-10/entity"

type ProductService interface {
	Create(payload entity.ProductRequest, userId uint) (entity.Product, error)
	Delete(product entity.Product)
	FindById(id uint) (entity.Product, error)
	Update(payload entity.ProductRequest, id uint, userId uint) (entity.Product, error)
}
