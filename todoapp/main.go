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
	Text      string `json:"text"`
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello derivco",
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

	r.GET("/todos/:itemid", func(c *gin.Context) {
		itemid := c.Param("itemid")
		item, err := query(itemid)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.JSON(200, item)
	})

	r.POST("/todos", func(c *gin.Context) {
		item := todoItem{}
		if err := c.Bind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		host, ok := os.LookupEnv("HOST")
		if !ok {
			host = "http://localhost:8080"
		}

		item.ID = uuid.New().String()
		item.URL = host + "/todos/" + item.ID

		err := insert(item)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.JSON(201, item)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func query(itemid string) (todoItem, error) {
	conn, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		conn = "postgres://postgres:mypassword@localhost:5432/postgres?pool_max_conns=10"
	}

	dbpool, err := pgxpool.Connect(context.Background(), conn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return todoItem{}, err
	}

	item := todoItem{}
	err = dbpool.QueryRow(context.Background(), "select id,title,url,completed,text from todo.todo where id=$1", itemid).Scan(&item.ID, &item.Title, &item.URL, &item.Completed, &item.Text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to query todo: %v\n", err)
		return todoItem{}, err
	}

	return item, nil
}

func insert(item todoItem) error {
	conn, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		conn = "postgres://postgres:mypassword@localhost:5432/postgres?pool_max_conns=10"
	}

	dbpool, err := pgxpool.Connect(context.Background(), conn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}

	defer dbpool.Close()

	_, err = dbpool.Exec(context.Background(), "insert into todo.todo values($1,$2,$3,$4,$5)", item.ID, item.Title, item.URL, item.Completed, item.Text)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert todo: %v\n", err)
		return err
	}

	return nil
}
