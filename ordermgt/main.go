package main

import (
	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func main() {
	r := gin.Default()

	r.GET("/health", healthCheck)

	r.Run(":9000")
}
