package main

import (
	"net/http"

	controllers "com.gcc/controller"
	"com.gcc/service"
	"github.com/gin-gonic/gin"
)

var (
	bookService    service.BookService        = service.New()
	bookController controllers.BookController = controllers.New(bookService)
)

//Dependency defines the dependencies of the app
type Dependency struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

//Health the health
type Health struct {
	Dependencies []Dependency `json:"dependencies"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/health", func(c *gin.Context) {

		dependencies := &Health{
			Dependencies: []Dependency{
				Dependency{Name: "one", Status: "pending"}}}

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
