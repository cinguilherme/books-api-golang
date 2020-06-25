package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
)

type Dependency struct {
	name string `json:"name"`
	status string `json:"status"`
}

type Health struct {
	dependencies []Dependency `json:"dependencies"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
	})

	r.GET("/health", func(c *gin.Context) {
		
		dependencies := &Health{
			dependencies: []Dependency{ 
				Dependency{ name: "one", status: "pending"}}}

		
		c.JSON(http.StatusOK, 
			gin.H{"data": "hello world", 
			"dependencies": dependencies})  
	})

	r.GET("/books", func(c *gin.Context) {
				
		c.JSON(http.StatusOK, 
			gin.H{"data": "books"})  
	})

	r.Run()
}