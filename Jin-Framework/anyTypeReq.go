package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person1 struct {
	Name string `form:"name" json:"name"`
	Age  string `form:"age" json:"age"`
}

func main() {
	router := gin.Default()

	router.Any("/", func(ctx *gin.Context) {
		var person Person1
		if err := ctx.ShouldBindQuery(&person); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"person":  person,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
		}
	})
	router.Run(":3000")
}
