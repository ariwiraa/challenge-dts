package routers

import (
	"challenge/controllers"

	"github.com/gin-gonic/gin"
)

func StartSever() *gin.Engine {
	router := gin.Default()
	api := router.Group("books")
	api.POST("", controllers.CreateBook)
	api.GET("", controllers.GetAllBooks)
	api.GET("/:id", controllers.GetBookById)
	api.PUT("/:id", controllers.UpdateBook)
	api.DELETE("/:id", controllers.DeleteBook)

	return router
}
