package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {
		var user User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		if user.Email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
			return
		} else if user.ID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
			return
		} else if user.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
			return
		} else if len(user.Hobbies) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Minimum 1 hobbies required"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"ID":      user.ID,
			"Name":    user.Name,
			"Email":   user.Email,
			"Hobbies": user.Hobbies,
			"Message": "User created successfully",
		})
	})

	// Start the server on port 3000
	router.Run(":3000")
}
