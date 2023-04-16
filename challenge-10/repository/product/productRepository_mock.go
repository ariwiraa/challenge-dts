package repository

import (
	"challange-10/entity"
	"log"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id uint) (entity.Product, error) {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		log.Println(nil)
	}

	product := arguments.Get(0).(entity.Product)

	return product, nil
}

func (repository *ProductRepositoryMock) Create(product entity.Product) (entity.Product, error) {
	panic("implement me")
}

func (repository *ProductRepositoryMock) Delete(product entity.Product) {
	panic("implement me")
}

func (repository *ProductRepositoryMock) Update(product entity.Product, id uint) (entity.Product, error) {
	panic("implement me")
}
