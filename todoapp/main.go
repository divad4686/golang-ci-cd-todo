package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Sum(a, b int) int {
	return a + b
}

type TodoItem struct {
	Id        string `json:"-"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Text      string `json:"text"`
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

	r.POST("/todos", func(c *gin.Context) {
		item := TodoItem{}
		if err := c.Bind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		item.Id = uuid.New().String()
		item.Url = "http://cicdexample.com/staging/todoapi/todos/" + item.Id
		c.Writer.Header().Add("Location", "/todos/"+item.Id)
		c.JSON(201, item)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
