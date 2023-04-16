package routers

import (
	"belajar-gin/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler handlers.BookHandler) *gin.Engine {
	router := gin.Default()
	api := router.Group("books")
	api.POST("", handler.PostBookHandler)
	api.GET("", handler.GetBooksHandler)
	api.GET("/:id", handler.GetBookByIdHandler)
	api.PUT("/:id", handler.PutBookHandler)
	api.DELETE("/:id", handler.DeleteBookHandler)

	return router
}
