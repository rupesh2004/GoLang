package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/",func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		nick :=ctx.DefaultPostForm("nick","anonymous")
		ctx.JSON(http.StatusOK,gin.H{
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":3000")
}