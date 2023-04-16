package handlers

import (
	"belajar-gin/entity"
	"belajar-gin/helper"
	"belajar-gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandlerImpl struct {
	BookService services.BookService
}

func NewBookHandler(bookService services.BookService) BookHandler {
	return &BookHandlerImpl{
		BookService: bookService,
	}
}

// deleteBookHandler implements BookHandler
func (service *BookHandlerImpl) DeleteBookHandler(ctx *gin.Context) {
	requestParams := ctx.Param("id")
	bookId, _ := strconv.Atoi(requestParams)

	service.BookService.DeleteBook(bookId)
	response := helper.APIResponse("delete succesfully", http.StatusOK, "success", nil)
	ctx.JSON(http.StatusOK, response)

}

// getBookByIdHandler implements BookHandler
func (service *BookHandlerImpl) GetBookByIdHandler(ctx *gin.Context) {
	requestParams := ctx.Param("id")
	bookId, _ := strconv.Atoi(requestParams)

	book, err := service.BookService.FindBookById(bookId)
	if err != nil {
		response := helper.APIResponse("failed get book id not found", http.StatusNotFound, "error", err)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Detail book", http.StatusOK, "success", book)
	ctx.JSON(http.StatusOK, response)
}

// getBooksHandler implements BookHandler
func (service *BookHandlerImpl) GetBooksHandler(ctx *gin.Context) {
	books, err := service.BookService.FindAllBooks()
	if err != nil {
		response := helper.APIResponse("failed get books", http.StatusBadRequest, "error", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of books", http.StatusOK, "success", books)
	ctx.JSON(http.StatusOK, response)

}

// postBookHandler implements BookHandler
func (handler *BookHandlerImpl) PostBookHandler(ctx *gin.Context) {
	var request entity.BookRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook, err := handler.BookService.AddBook(request)

	if err != nil {
		response := helper.APIResponse("failed created book", http.StatusBadRequest, "error", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success created a book", http.StatusCreated, "success", newBook)
	ctx.JSON(http.StatusCreated, response)

}

// putBookHandler implements BookHandler
func (service *BookHandlerImpl) PutBookHandler(ctx *gin.Context) {
	requestParams := ctx.Param("id")
	var request entity.BookRequest
	bookId, _ := strconv.Atoi(requestParams)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedBook, err := service.BookService.UpdateBook(request, bookId)
	if err != nil {
		response := helper.APIResponse("failed update book", http.StatusBadRequest, "error", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("success created a book", http.StatusCreated, "success", updatedBook)
	ctx.JSON(http.StatusOK, response)
}
