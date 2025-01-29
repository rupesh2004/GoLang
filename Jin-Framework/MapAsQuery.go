// http://localhost:3000/?ids[a]=1&ids[b]=2&names[name1]=rupesh&names[name2]=monal

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ids :=ctx.QueryMap("ids")
		name :=ctx.QueryMap("names")

		ctx.JSON(http.StatusOK,gin.H{
			"ids":ids,
			"names":name,
		})
	})
	router.Run(":3000")
}
