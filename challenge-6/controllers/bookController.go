package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.Id = fmt.Sprintf("book-%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, "Created")
}

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, BookDatas)
}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("id") // mengambil request parameter
	condition := false

	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookId == book.Id {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].Id = bookId
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

func GetBookById(ctx *gin.Context) {
	bookId := ctx.Param("id") // mengambil request parameter
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookId == book.Id {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id") // mengambil request parameter
	condition := false

	var bookIndex int
	for i, book := range BookDatas {
		if bookId == book.Id {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, "Deleted")
}
