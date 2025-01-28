package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Serve static files from the "html" folder
	r.Static("/html", "./html")

	// Serve the static index.html file from the root route
	r.GET("/", func(c *gin.Context) {
		c.File("./html/index.html") // Serve the external index.html file directly
	})

	// Run the server
	r.Run(":8080")
}
