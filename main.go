package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong2",
		})
	})

	return r
}

func main() {
	r := router()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
