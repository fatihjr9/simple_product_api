package main

import (
	"product/controllers"
	"product/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Call DB
	models.ConnectDatabase()

	// Router
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// All post routes
	router.GET("/api/product", controllers.GetPosts)
	router.POST("/api/product", controllers.StorePost)
	router.GET("/api/product/:id", controllers.DetailPost)
	router.PUT("/api/product/:id", controllers.UpdatePost)
	router.DELETE("/api/product/:id", controllers.DeletePost)

	router.Run(":3000")
}
