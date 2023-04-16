package main

import (
	"belajar-gin/app"
	"belajar-gin/handlers"
	"belajar-gin/repository"
	"belajar-gin/routers"
	"belajar-gin/services"
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
