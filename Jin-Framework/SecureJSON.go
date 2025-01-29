package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/",func(ctx *gin.Context) {
		name := []string{"rupesh","monal"}
		ctx.SecureJSON(200,name)
	})
	router.Run(":3000")
}