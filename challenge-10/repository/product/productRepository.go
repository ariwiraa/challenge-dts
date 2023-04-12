package repository

import "challange-10/entity"

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	FindById(id uint) (entity.Product, error)
	Update(product entity.Product, id uint) (entity.Product, error)
	Delete(product entity.Product)
}
