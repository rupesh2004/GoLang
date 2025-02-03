package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Students struct {
	Name   string `json:"name" binding:"required" validate:"required"`
	Age    string `json:"age" binding:"required"`
	Grade  string `json:"grade" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

const uri = "mongodb+srv://bhosalerupesh67:yjk7k48j3k4R2Zde@studentsdata.151co.mongodb.net/?retryWrites=true&w=majority&appName=studentsData"

var mongoClient *mongo.Client
var studentCollection *mongo.Collection

func init() {
	if error := connect_to_mongo(); error != nil {
		fmt.Println("could not connect to mongo")
	} else {
		fmt.Println("connected to mongo")
	}
	studentCollection = mongoClient.Database("studentsData").Collection("students")
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		var students []Students
		cursor, err := studentCollection.Find(context.TODO(), bson.D{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not fetch students"})
			return
		}
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			var student Students
			if err = cursor.Decode(&student); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "could not decode student"})
				return
			}
			students = append(students, student)
		}
		if err = cursor.Err(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not fetch students"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"students": students})
	})

	router.POST("/addStudents", func(ctx *gin.Context) {
		var addStudent Students
		if error := ctx.ShouldBind(&addStudent); error != nil {
			ctx.JSON(400, gin.H{"error": error.Error()})
			return
		}
		insertedResult, err := studentCollection.InsertOne(context.TODO(), addStudent)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not insert student"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "student inserted", "id": insertedResult.InsertedID})
	})

	router.DELETE("/deleteStudent/:email", func(ctx *gin.Context) {
		email := ctx.Param("email")
		_, err := studentCollection.DeleteOne(context.TODO(), bson.D{{"email", email}})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not delete student"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "student deleted"})
	})

	router.PUT("/updateStudent/:email", func(ctx *gin.Context) {
		email := ctx.Param("email")
		var updateStudent Students
		if error := ctx.ShouldBind(&updateStudent); error != nil {
			ctx.JSON(400, gin.H{"error": error.Error()})
			return
		}
		_, error := studentCollection.UpdateOne(context.TODO(), bson.D{{"email", email}}, bson.D{{"$set", updateStudent}})
		if error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not update student"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "student updated"})
	})

	router.Run(":3000")
}

func connect_to_mongo() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, error := mongo.Connect(context.TODO(), opts)
	if error != nil {
		return error
	}
	error = client.Ping(context.TODO(), nil)
	mongoClient = client
	return error

}
