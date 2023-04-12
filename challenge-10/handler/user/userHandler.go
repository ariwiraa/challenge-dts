package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	PostUserHandler(ctx *gin.Context)
	PostUserLoginHandler(ctx *gin.Context)
}
