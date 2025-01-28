package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/",func(ctx *gin.Context) {
		data := gin.H {
			"lang":"GO语言",
			"tag" : "<br>",
		}
		ctx.JSON(200,gin.H{
			"data":data,
		})
		ctx.AsciiJSON(http.StatusOK,data)
	})
	router.Run(":3000")
}