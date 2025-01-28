package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `uri:"name"`
	Age  string `uri:"age"`
}

func main() {
	router := gin.Default()

	// Define the route to bind the parameters
	router.GET("/:age/:name", func(ctx *gin.Context) {
		var p Person
		// Use ShouldBindUri for binding URL parameters
		if err := ctx.ShouldBindUri(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"name": p.Name, "age": p.Age})
		}
	})

	router.Run(":3000")
}
