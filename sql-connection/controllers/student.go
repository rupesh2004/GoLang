package controllers

import (
	"database/sql"
	"net/http"
	"sql-connection/config"
	"sql-connection/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func InitController() {
	if config.DB == nil {
		panic("Database not connected")
	}
}

func FormatDate(birthDate string) (string, error) {
	parsedDate, err := time.Parse("02-01-2006", birthDate)
	if err != nil {
		return "", err
	}
	return parsedDate.Format("2006-01-02"), nil
}

func CheckStudentExists(prn int) (bool, error) {
	var existingPRN int
	err := config.DB.QueryRow("SELECT prn FROM stud WHERE prn = ?", prn).Scan(&existingPRN)
	if err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	}
	return false, err
}

func AddStudent(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	exists, err := CheckStudentExists(student.PRN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if exists {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Student already present"})
		return
	}

	formattedDate, err := FormatDate(student.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use DD-MM-YYYY"})
		return
	}
	student.BirthDate = formattedDate

	_, err = config.DB.Exec("INSERT INTO stud (name, prn, mobile_no, birth_date) VALUES (?, ?, ?, ?)",
		student.Name, student.PRN, student.MobileNo, student.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student added successfully!"})
}

func GetStudents(ctx *gin.Context) {
	rows, err := config.DB.Query("SELECT name, prn, mobile_no, birth_date FROM stud")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.Name, &student.PRN, &student.MobileNo, &student.BirthDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		students = append(students, student)
	}

	ctx.JSON(http.StatusOK, students)
}

func UpdateStudent(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	exists, err := CheckStudentExists(student.PRN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	formattedDate, err := FormatDate(student.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use DD-MM-YYYY"})
		return
	}
	student.BirthDate = formattedDate

	_, err = config.DB.Exec("UPDATE stud SET name = ?, mobile_no = ?, birth_date = ? WHERE prn = ?",
		student.Name, student.MobileNo, student.BirthDate, student.PRN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student updated successfully!"})
}

func DeleteStudent(ctx *gin.Context) {
	prnStr := ctx.Param("prn")
	prn, err := strconv.Atoi(prnStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PRN format"})
		return
	}
	exists, err := CheckStudentExists(prn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}
	_,err = config.DB.Exec("delete from stud where prn = ?", prn)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"Student deleted successfully"})
}

func GetDataByPRN(ctx *gin.Context){
	prnStr := ctx.Param("prn")
	prn, err := strconv.Atoi(prnStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PRN format"})
		return
	}
	exists,err := CheckStudentExists(prn)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}
	var student models.Student
	err = config.DB.QueryRow("select * from stud where prn = ?", prn).Scan(&student.Name, &student.PRN, &student.MobileNo, &student.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"student": student})
}
