package service

import (
	product "challange-10/repository/product"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &product.ProductRepositoryMock{
	Mock: mock.Mock{},
}

// var userRepository = &repository.UserRepositoryMock{
// 	Mock: mock.Mock{},
// }

var productService = ProductServiceImpl{
	ProductRepository: productRepository,
	// UserRepository:    userRepository,
}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", uint(1)).Return(nil)

	product, err := productService.FindById(1)
	fmt.Println("ini product", product)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found")
}
