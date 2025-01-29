package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Email    string `form:"Email"`
	Password string `form:"Password"`
}

func main() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {
		var loginForm LoginForm
		if err := ctx.ShouldBind(&loginForm); err == nil {
			if loginForm.Email == "r@gmail.com" && loginForm.Password == "1234" {
				ctx.JSON(200, gin.H{"message": "Login Success"})
			} else {
				ctx.JSON(401, gin.H{"message": "Login Failed"})

			}
		}
	})
	router.Run(":3000")
}
