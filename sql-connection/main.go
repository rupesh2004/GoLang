package main

import (
	"sql-connection/config"
	"sql-connection/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()
	routes.UseRoutes(router)
	router.Run(":3000")
}
