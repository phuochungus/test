package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	sqlcGen "tutorial.sqlc.dev/app/db/generated"
)

func connectDB() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:123123123@localhost:5432/simple_library")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	q := sqlcGen.New(conn)

	author, err := q.GetAuthor(context.Background(), 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(author.Name)
}

func main() {
	connectDB()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
