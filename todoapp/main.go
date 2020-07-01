package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Sum a number
func Sum(a, b int) int {
	return a + b
}

type todoItem struct {
	ID        string `json:"-"`
	Title     string `json:"title"`
	URL       string `json:"url"`
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
		item := todoItem{}
		if err := c.Bind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		item.ID = uuid.New().String()
		item.URL = "http://cicdexample.com/staging/todoapi/todos/" + item.ID

		err := insert(item)
		if err != nil {
			c.JSON(500, err)
		}
		c.JSON(201, item)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func insert(item todoItem) error {
	conn, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		conn = "postgres://postgres:mypassword@localhost:5432/postgres?pool_max_conns=10&currentSchema=todo"
	}

	dbpool, err := pgxpool.Connect(context.Background(), conn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}

	defer dbpool.Close()

	_, err = dbpool.Exec(context.Background(), "insert into todo values($1,$2,$3,$4,$5,$6)", item.ID, item.Title, item.URL, item.Completed, item.Order, item.Text)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert todo: %v\n", err)
		return err
	}

	return nil
}
