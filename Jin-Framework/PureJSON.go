package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"html":"<b>Hello</b>",
		})
	})
	router.GET("/purejson",func(ctx *gin.Context) {
		ctx.PureJSON(http.StatusOK,gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	router.Run(":3000")
}
