package main

import (
	"challange-10/app"
	handlerProduct "challange-10/handler/product"
	handler "challange-10/handler/user"
	repositoryProduct "challange-10/repository/product"
	repository "challange-10/repository/user"
	"challange-10/routes"
	serviceProduct "challange-10/services/product"
	service "challange-10/services/user"
)

func main() {
	db := app.InitializeDatabase()
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := handler.NewUserHandlerImpl(userService)

	productRepository := repositoryProduct.NewProductRepository(db)
	productService := serviceProduct.NewProductService(productRepository, userRepository)
	productHandler := handlerProduct.NewProductHandler(productService)
	router := routes.Route(userHandler, productHandler)

	router.Run(":8000")
}
