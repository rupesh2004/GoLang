package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/test"
		router.HandleContext(ctx)
	})
	router.GET("/test", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/test2"
		router.HandleContext(ctx)
	})
	router.GET("/test2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"hello": "world"})
	})
	router.Run(":3000")
}
