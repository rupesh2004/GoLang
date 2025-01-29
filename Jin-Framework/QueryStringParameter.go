package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/",func(ctx *gin.Context) {
		id :=ctx.DefaultQuery("id","0")
		name :=ctx.Query("name")
		if name ==""{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error": "name is required",
			})
			return
		}

		ctx.JSON(200,gin.H{"id":id,"name":name})
	})
	router.Run(":3000")
}