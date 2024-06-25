package book

import (
	"go-web-api/struct/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostBookHandler(c *gin.Context) {

	var book book.Book

	err := c.BindJSON(&book)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"errors": err.Error(),
		// })

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors.Errors(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":  book.Title,
		"author": book.Author,
	})
}

func GetBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Buku Bahagia",
	})
}

func GetBookByIdHandler(c *gin.Context) {
	id := c.Param("id")

	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"title": "Buku Bahagia dengan id " + id + " dan query " + title,
	})
}
