package handlers

import "github.com/gin-gonic/gin"

type BookHandler interface {
	PostBookHandler(ctx *gin.Context)
	GetBooksHandler(ctx *gin.Context)
	GetBookByIdHandler(ctx *gin.Context)
	PutBookHandler(ctx *gin.Context)
	DeleteBookHandler(ctx *gin.Context)
}
