package service

import (
	"challange-10/entity"

	repository "challange-10/repository/product"
	repositoryUser "challange-10/repository/user"
	"log"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	UserRepository    repositoryUser.UserRepository
}

// Delete implements ProductService
func (service *ProductServiceImpl) Delete(product entity.Product) {
	product, err := service.ProductRepository.FindById(product.ID)
	if err != nil {
		panic(err)
	}

	service.ProductRepository.Delete(product)
}

// FindById implements ProductService
func (service *ProductServiceImpl) FindById(id uint) (entity.Product, error) {
	product, err := service.ProductRepository.FindById(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

// Update implements ProductService
func (service *ProductServiceImpl) Update(payload entity.ProductRequest, id uint, userId uint) (entity.Product, error) {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		log.Println("tidak ada user dengan id tersebut")
	}

	product, err := service.ProductRepository.FindById(id)
	if err != nil {
		return product, err
	}

	product.Title = payload.Title
	product.Description = payload.Description
	product.UserId = userId
	product.User = &user

	product, err = service.ProductRepository.Update(product, id)
	if err != nil {
		return product, err
	}

	return product, nil
}

// Create implements ProductService
func (service *ProductServiceImpl) Create(payload entity.ProductRequest, userId uint) (entity.Product, error) {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		log.Println("tidak ada user dengan id tersebut")
	}

	product := entity.Product{
		Title:       payload.Title,
		Description: payload.Description,
		UserId:      userId,
		User:        &user,
	}

	newProduct, err := service.ProductRepository.Create(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func NewProductService(productRepository repository.ProductRepository, userRepository repositoryUser.UserRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository,
		UserRepository: userRepository}
}
