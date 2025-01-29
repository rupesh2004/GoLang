package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		cookie,err :=c.Cookie("rupesh")
		if err != nil {
			c.SetCookie("rupesh","hello from rupesh",3600,"/","localhost",false,true)
			c.JSON(200, gin.H{"message": "Cookie set"})
			return

		}
		c.JSON(200, gin.H{"message": cookie})
	})
	router.Run(":3000")
}

