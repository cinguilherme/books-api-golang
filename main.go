package main

import (
	"net/http"

	controllers "com.gcc/controller"
	"com.gcc/entity"
	"com.gcc/service"
	"github.com/gin-gonic/gin"
)

var (
	bookService    service.BookService        = service.New()
	bookController controllers.BookController = controllers.New(bookService)
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/health", func(c *gin.Context) {

		dependencies := &entity.Health{
			Dependencies: []entity.Dependency{
				entity.Dependency{Name: "one", Status: "pending"}}}

		c.JSON(http.StatusOK,
			gin.H{"data": "hello world",
				"dependencies": dependencies})
	})

	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, bookController.FindAll())
	})
	r.POST("/books", func(c *gin.Context) {
		c.JSON(http.StatusCreated,
			bookController.Save(c))
	})

	r.Run(":3000")
}
