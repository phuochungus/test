package main

import (
	"root/src/authors"
	"root/src/pool"

	"github.com/gin-gonic/gin"
)

func main() {
	pool.CreatePool()
	defer pool.DestroyPool()
	r := gin.Default()
	authors.CreateController(r)
	r.Run()
}
