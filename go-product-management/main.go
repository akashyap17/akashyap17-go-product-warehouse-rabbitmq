package main

import (
	"github.com/akashyap17/go-product-management/db"
	"github.com/akashyap17/go-product-management/rabbitmq"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	db.InitDB()

	// Initialize RabbitMQ
	rabbitmq.InitRabbitMQ()

	// Gin Router
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	// Start server
	r.Run(":8080")
}
