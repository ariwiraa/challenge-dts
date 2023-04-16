package main

import (
	"gin-gorm/app"
	"gin-gorm/handlers"
	"gin-gorm/repository"
	"gin-gorm/routers"
	"gin-gorm/services"
)

func main() {
	var port = ":4000"

	db := app.InitializeDatabase()
	bookRepository := repository.NewBookRepository(db)
	bookService := services.NewBookServiceImpl(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)
	router := routers.NewRouter(bookHandler)

	router.Run(port)
}
