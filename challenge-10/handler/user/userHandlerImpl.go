package handler

import (
	"challange-10/entity"
	"challange-10/helpers"
	service "challange-10/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct {
	UserService service.UserService
}

// PostUserLoginHandler implements UserHandler
func (handler *UserHandlerImpl) PostUserLoginHandler(ctx *gin.Context) {
	var payload entity.UserLogin

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	loggedInUser, err := handler.UserService.Login(payload)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token := helpers.GenerateToken(loggedInUser.ID, loggedInUser.Email, loggedInUser.IsAdmin)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})

}

// PostUserHandler implements UserHandler
func (handler *UserHandlerImpl) PostUserHandler(ctx *gin.Context) {
	var request entity.UserRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newUser, err := handler.UserService.AddUser(request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       newUser.ID,
		"email":    newUser.Email,
		"fullName": newUser.FullName,
	})
}

func NewUserHandlerImpl(userService service.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}
