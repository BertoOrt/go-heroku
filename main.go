package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"this": "works"})
	})
	r.Run(":" + port)
}
