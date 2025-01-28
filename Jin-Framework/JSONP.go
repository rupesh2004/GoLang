package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/",func(ctx *gin.Context) {
		callback := ctx.DefaultQuery("callback","")
		data := gin.H{
			"email":"r@gmail.com",
			"password":"123456",
		}
		if callback!=""{
			ctx.JSONP(http.StatusOK,gin.H{
				"message":"using json padding ",
				"data":data,
			})
		}else {
			ctx.JSON(http.StatusOK,gin.H{
				"message":"using simple json",
				"data":data,
			})
		}

	})
	router.Run(":3000")
}
