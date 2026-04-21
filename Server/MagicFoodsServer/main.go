package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hello, go World!")
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicFoodRecomendation!")
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
