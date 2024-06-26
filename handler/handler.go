package handler

import (
	"go-web-api/book"
	"go-web-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookHandler(ctx *gin.Context) {
	db := database.DatabaseInstance()

	// instance books
	books := []book.Book{}

	// find and assign data to books
	db.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func CreateBookHandler(ctx *gin.Context) {
	db := database.DatabaseInstance()

	db.Create(&book.Book{
		ID:          2,
		Title:       "Buku Ramah Lingkungan",
		Author:      "Reza",
		Description: "lorem insum",
	})

	// find book data
	// data := book.Book{}
	// db.Where("id = ?", 2).First(&data)
	// fmt.Println(data.Title)

	println("sukes buat create data buku")
}

func GetByIdBookHandler(ctx *gin.Context) {
	db := database.DatabaseInstance()

	id := ctx.Param("id")

	// instance book
	book := book.Book{}

	// find and assign data to books
	db.Where("id = ?", id).First(&book)

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func UpdateBookHandler(ctx *gin.Context) {
	db := database.DatabaseInstance()

	id := ctx.Param("id")
	// jika ingin mengambil dari formData request
	_title := ctx.PostForm("title")

	//
	bookData := book.Book{}

	// find old data
	db.Where("id = ?", id).First(&bookData)

	// set and update old data
	bookData.Title = _title
	db.Save(bookData)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookData,
	})
}

func DeleteBookHandler(ctx *gin.Context) {
	db := database.DatabaseInstance()

	id := ctx.Param("id")

	book := book.Book{}

	// delete data
	db.Where("id = ?", id).Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "sukses hapus data",
		"data":    book,
	})
}
