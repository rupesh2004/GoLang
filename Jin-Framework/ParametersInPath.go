package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/:id/:name",func(ctx *gin.Context) {
		id := ctx.Param("id")
		name := ctx.Param("name")
		ctx.JSON(200, gin.H {
			"id": id,
			"name": name,
		})
	})
	router.GET("/testing/:id/*action",func(ctx *gin.Context) {
		action := ctx.Param("action")
		id := ctx.Param("id")
		message := id+"hello"+action
		ctx.JSON(200, gin.H {
			"action": message,
		})
	})
	router.Run(":3000")
}
