package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// gin.DisableConsoleColor()
	gin.ForceConsoleColor()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome to gin",
		})
	})
	router.Run(":3000")
}
