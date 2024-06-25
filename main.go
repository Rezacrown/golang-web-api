package main

import (
	book "go-web-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/book", book.GetBookHandler)
	r.GET("/book/:id", book.GetBookByIdHandler)

	r.POST("/book", book.PostBookHandler)

	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080
}
