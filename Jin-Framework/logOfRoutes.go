package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	routes := gin.Default()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	routes.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "get route"})
	})
	routes.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "post route"})
	})
	routes.Run(":3000")
}
