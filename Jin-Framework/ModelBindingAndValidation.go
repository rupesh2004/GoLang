package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email string `form="email" json:"email" xml:"email" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
func main() {
	router := gin.Default()

	router.POST("/loginJSON",func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err!=nil{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"message": "Unauthorized",
			})
		}

		if user.Email =="rupesh@gmail.com" && user.Password =="rupesh@123" {
			ctx.JSON(http.StatusOK,gin.H{
				"message": "Login Success",
			})
		}else{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"message": "Invalid Credentials",
			})
		}
	})

	router.POST("loginXML",func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindXML(&user); err!=nil{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"message": "Unauthorized",
			})
		}
		if user.Email =="rupesh@gmail.com" && user.Password =="rupesh@123" {
			ctx.JSON(http.StatusOK,gin.H{
				"message": "Login Success",
			})
		}else {
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"message": "Invalid Credentials",
			})
		}
	})
	
	router.Run(":3000")
}