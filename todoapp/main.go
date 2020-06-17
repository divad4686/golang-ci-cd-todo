package main

import "github.com/gin-gonic/gin"

func Sum(a, b int) int {
	return a + b
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/sum", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": Sum(2, 2),
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "fine",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
