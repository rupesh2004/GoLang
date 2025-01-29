package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		name :=c.PostForm("name")
		message := c.PostForm("message")
		c.JSON(200, gin.H{
			"id": id,
			"name": name,
			"message": message,
			"page":page,
		})
	})
	router.Run(":3000")
}