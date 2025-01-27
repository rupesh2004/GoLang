package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:id", func(ctx *gin.Context) {
		var id = ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	})
	router.POST("/me",func(ctx *gin.Context) {
		type Person struct{
			Email string `json:"email" binding:"required"`
			Password string `json:password " binding:"required"`
		}
		var p Person
		ctx.BindJSON(&p)
		ctx.JSON(http.StatusOK,gin.H{
			"email": p.Email,
			"password": p.Password,
		})
	})

	router.Run(":3000")
}
