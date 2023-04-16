package routers

import (
	"gin-gorm/handlers"

	// _ "gin-gorm/docs" // docs is genereated by Swag CLI, you have to import it.

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bok API
// @version 1.0
// @description this is a sample service for managing cars
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
func NewRouter(handler handlers.BookHandler) *gin.Engine {
	router := gin.Default()
	api := router.Group("books")
	api.POST("", handler.PostBookHandler)
	api.GET("", handler.GetBooksHandler)
	api.GET("/:id", handler.GetBookByIdHandler)
	api.PUT("/:id", handler.PutBookHandler)
	api.DELETE("/:id", handler.DeleteBookHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
