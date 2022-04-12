package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.UseH2C = true

	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 13})
	})

	engine.Run(":8088")
}
