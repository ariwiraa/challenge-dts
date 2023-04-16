package handlers

import (
	"gin-gorm/entity"
	"gin-gorm/services"
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

// DeleteBooks godoc
// @Summary Delete book identified by the given id
// @Description Delete the book corresponding to the input Id
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "ID of the book to be deleted"
// @Success 204 "No Content"
// @Router /books/{id} [delete]
// deleteBookHandler implements BookHandler
func (service *BookHandlerImpl) DeleteBookHandler(ctx *gin.Context) {
	var book entity.Book
	requestParams := ctx.Param("id")
	book.Id, _ = strconv.Atoi(requestParams)

	service.BookService.DeleteBook(book)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete succesfully",
	})

}

// GetBookByIdHandler godoc
// @Summary Get Details for a given id
// @Description Get details of book corresponding is the input Id
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "ID of the books"
// @Success 200 {object} entity.Book
// @Router /books/{id} [get]
// getBookByIdHandler implements BookHandler
func (service *BookHandlerImpl) GetBookByIdHandler(ctx *gin.Context) {
	requestParams := ctx.Param("id")
	bookId, _ := strconv.Atoi(requestParams)

	book, err := service.BookService.FindBookById(bookId)
	if err != nil {

		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// GetBooksHandler godoc
// @Summary Get Details
// @Description Get details of all books
// @Tags book
// @Accept json
// @Produce json
// @Success 200 {object} entity.Book
// @Router /books [get]
// getBooksHandler implements BookHandler
func (service *BookHandlerImpl) GetBooksHandler(ctx *gin.Context) {
	books, err := service.BookService.FindAllBooks()
	if err != nil {

		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, books)

}

// CreateBooks godoc
// @Summary Post Details
// @Description Post details of book
// @Tags book
// @Accept json
// @Produce json
// @Param entity.Books body entity.Book true "create book"
// @Success 200 {object} entity.Book
// @Router /books [post]
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
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, newBook)

}

// UpdateBooks godoc
// @Summary Update book identified by the given id
// @Description Update the book corresponding is the input Id
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "ID of the book to be updated"
// @Success 200 {object} entity.Book
// @Router /books/{id} [put]
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

		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
}
