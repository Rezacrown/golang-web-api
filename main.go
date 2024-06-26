package main

import (
	"go-web-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//
	v1 := r.Group("/v1")

	v1.GET("/book", handler.GetBookHandler)
	v1.GET("/book/:id", handler.GetByIdBookHandler)

	v1.POST("/book", handler.CreateBookHandler)

	v1.PUT("/book/:id", handler.UpdateBookHandler)

	v1.DELETE("/book/:id", handler.DeleteBookHandler)

	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080
}
