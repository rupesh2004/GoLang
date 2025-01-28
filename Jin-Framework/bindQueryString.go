// http://localhost:3000/?name=monal&age=22&prn=2253010
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type stud struct {
	Name string `form:"name" json:"name"`
	Age  string `form:"age" json:"age"`
	PRN  string `form:"prn" json:"prn"`
}

func student(ctx *gin.Context) {
	var s stud
	if err := ctx.ShouldBind(&s); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Name": s.Name,
			"Age":  s.Age,
			"PRN":  s.PRN,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid input"})
	}
}
func main() {
	router := gin.Default()
	router.GET("/", student)
	router.Run(":3000")
}

// package main

// import "github.com/gin-gonic/gin"

// type stud struct {
// 	Name string `json:"name"`
// 	Age  string `json:"age"`
// 	PRN  string `json:"prn"`
// }

// func student(ctx *gin.Context) {
// 	// Manually fetch query parameters
// 	name := ctx.DefaultQuery("name", "")
// 	age := ctx.DefaultQuery("age", "")
// 	prn := ctx.DefaultQuery("prn", "")

// 	// Create a new struct with the values from the query parameters
// 	s := stud{
// 		Name: name,
// 		Age:  age,
// 		PRN:  prn,
// 	}

// 	// Return the response with the populated struct
// 	ctx.JSON(200, gin.H{
// 		"name": s.Name,
// 		"age":  s.Age,
// 		"PRN":  s.PRN,
// 	})
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/", student)
// 	router.Run(":3000")
// }
