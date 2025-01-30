package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "file not uploaded"})
			return
		}
		dst := filepath.Join("uploads", file.Filename)
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to save file",
			})
			return
		}
		ctx.JSON(200, gin.H{"message": "file uploaded successfully", "file name": file.Filename})

	})
	router.Run(":3000")
}
