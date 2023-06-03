package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	InitAppRouting(r)

	r.Run()
}
