package handler

import (
	"challange-10/entity"
	service "challange-10/services/product"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductHandlerImpl struct {
	ProductService service.ProductService
}

// DeleteProductHandler implements ProductHandler
func (handler *ProductHandlerImpl) DeleteProductHandler(ctx *gin.Context) {
	var product entity.Product
	requestParams := ctx.Param("id")
	productId, _ := strconv.Atoi(requestParams)
	product.ID = uint(productId)

	handler.ProductService.Delete(product)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete succesfully",
	})
}

// GetProductByIdHandler implements ProductHandler
func (handler *ProductHandlerImpl) GetProductByIdHandler(ctx *gin.Context) {
	requestParams := ctx.Param("id")
	productId, _ := strconv.Atoi(requestParams)

	product, err := handler.ProductService.FindById(uint(productId))
	if err != nil {

		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// PutProductHandler implements ProductHandler
func (handler *ProductHandlerImpl) PutProductHandler(ctx *gin.Context) {
	requestParams := ctx.Param("id")
	var request entity.ProductRequest
	productId, _ := strconv.Atoi(requestParams)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedProduct, err := handler.ProductService.Update(request, uint(productId), userID)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

// PostProductHandler implements ProductHandler
func (handler *ProductHandlerImpl) PostProductHandler(ctx *gin.Context) {
	var payload entity.ProductRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product, err := handler.ProductService.Create(payload, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)

}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &ProductHandlerImpl{ProductService: productService}
}
