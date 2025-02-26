package routes

import (
	"sql-connection/controllers"

	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine){
	// Define routes here
	controllers.InitController()
	r.POST("/addStudent",controllers.AddStudent)
	r.GET("/getStudents",controllers.GetStudents)
	r.PUT("/updateStudent",controllers.UpdateStudent)
	r.DELETE("/deleteStudent/:prn",controllers.DeleteStudent)
	r.GET("/getByPRN/:prn",controllers.GetDataByPRN)
}
