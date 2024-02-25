package main

import (
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, strconv.Itoa(rand.Intn(100)))
	})

	r.Run(":3030")
}
