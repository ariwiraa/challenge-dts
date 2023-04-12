package handler

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	PostProductHandler(ctx *gin.Context)
	GetProductByIdHandler(ctx *gin.Context)
	PutProductHandler(ctx *gin.Context)
	DeleteProductHandler(ctx *gin.Context)
}
