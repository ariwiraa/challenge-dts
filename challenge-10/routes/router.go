package routes

import (
	handlerProduct "challange-10/handler/product"
	handler "challange-10/handler/user"
	"challange-10/middlewares"

	"github.com/gin-gonic/gin"
)

// type Router struct {
// 	UserHandler *handler.UserHandler
// }

// func NewRouter(userHandler *handler.UserHandler) *Router {
// 	return &Router{UserHandler: userHandler}
// }

func Route(userHandler handler.UserHandler, productHandler handlerProduct.ProductHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", userHandler.PostUserHandler)
	router.POST("/signin", userHandler.PostUserLoginHandler)

	productRouter := router.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("", productHandler.PostProductHandler)
		productRouter.GET("/:id", middlewares.ProductAuthorization(), productHandler.GetProductByIdHandler)
		productRouter.PUT("/:id", middlewares.AuthAdminMiddlerware(), productHandler.PutProductHandler)
		productRouter.DELETE("/:id", middlewares.AuthAdminMiddlerware(), productHandler.DeleteProductHandler)

	}

	return router
}
