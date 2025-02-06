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
	ID     string `bson:"_id" json:"id"`
	Name   string `json:"name" binding:"required" validate:"required"`
	Age    string `json:"age" binding:"required"`
	Marks  []int  `json:"marks" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

const uri = "mongodb+srv://bhosalerupesh67:6LY0A6kY0mwEihmy@mongo-aggregate.jntyj.mongodb.net/?retryWrites=true&w=majority&appName=Mongo-Aggregate"

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

func createStudent(c *gin.Context) {
	var student Students
	if error := c.ShouldBindJSON(&student); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	insertResult, error := studentCollection.InsertOne(context.TODO(), student)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not insert student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "student created", "id": insertResult.InsertedID})

}

func addFields(c *gin.Context) {
	pipeLine := mongo.Pipeline{
		{{"$addFields", bson.D{{"totalMarks", bson.D{{"$sum", "$marks"}}}}}},
		{{"$addFields", bson.D{{"averageMarks", bson.D{{"$divide", bson.A{"$totalMarks", bson.D{{"$size", "$marks"}}}}}}}}},
	}

	cursor, error := studentCollection.Aggregate(context.TODO(), pipeLine)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not add fields"})
		return
	}
	var students []bson.M
	if error = cursor.All(context.TODO(), &students); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not add fields"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func count(c *gin.Context) {
	pipeLine := mongo.Pipeline{
		{{"$addFields", bson.D{{"intAge", bson.D{{"$toInt", "$age"}}}}}},
		{{"$match", bson.D{{"intAge", bson.D{{"$gte", 26}}}}}},
		{{"$count", "totalStudents"}},
	}
	cursor, error := studentCollection.Aggregate(context.TODO(), pipeLine)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count students"})
		return
	}
	var students []bson.M
	if error = cursor.All(context.TODO(), &students); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func group(c *gin.Context) {
	pipeLine := mongo.Pipeline{
		bson.D{
			{"$group", bson.D{
				{"_id", "$age"},
				{"names", bson.D{{"$push", "$name"}}},
				{"totalStudents", bson.D{{"$sum", 1}}},
			}},
		},
	}
	cursor, error := studentCollection.Aggregate(context.TODO(), pipeLine)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	var students []bson.M
	if error = cursor.All(context.TODO(), &students); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": error.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func match(ctx *gin.Context) {
	pileLine := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"age", bson.D{{"$gt", "25"}}},
			}},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", "$age"},
				{"name", bson.D{{"$push", "$name"}}},
				{"totalStudents", bson.D{{"$sum", 1}}},
			}},
		},
	}
	cursor, error := studentCollection.Aggregate(context.TODO(), pileLine)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not count students"})
		return
	}
	var student []bson.M
	if error = cursor.All(context.TODO(), &student); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, student)

}

func sort(ctx *gin.Context) {
	pipeLine := mongo.Pipeline{
		bson.D{
			{"$sort", bson.D{{"name", 1}}},
		},
	}
	cursor, error := studentCollection.Aggregate(context.TODO(), pipeLine)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	var student []bson.M
	if error = cursor.All(context.TODO(), &student); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	}
	ctx.JSON(http.StatusOK, student)
}

func lookup(ctx *gin.Context) {
	pipeLine := mongo.Pipeline{
		bson.D{
			{"$lookup", bson.D{
				{"from", "courses"},
				{"localField", "course_id"},
				{"foreignField", "courseID"},
				{"as", "courseDetails"},
			}},
		},
	}
	cursor, error := studentCollection.Aggregate(context.TODO(), pipeLine)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"ereor": error})
		return
	}
	var student []bson.M
	if error = cursor.All(context.TODO(), &student); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, student)
}

func mergeDoc(ctx *gin.Context) {
	pipeLine := mongo.Pipeline{
		// Add fields (totalMarks and averageMarks)
		bson.D{
			{"$addFields", bson.D{
				// Sum the marks array
				{"totalMarks", bson.D{{"$sum", "$marks"}}},
				// Calculate average marks by dividing totalMarks by the size of the marks array
				{"averageMarks", bson.D{{"$divide", bson.A{"$totalMarks", bson.D{{"$size", "$marks"}}}}}},
			}},
		},
		// Merge the results back into the students collection
		bson.D{
			{"$merge", bson.D{
				{"into", "students"},
				{"whenMatched", "replace"},
				{"whenNotMatched", "insert"},
			}},
		},
	}

	cursor, err := studentCollection.Aggregate(context.TODO(), pipeLine)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var students []bson.M
	if err = cursor.All(context.TODO(), &students); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func out(ctx *gin.Context) {
	pipeLine := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"age", bson.D{{"$gt", "20"}}},
			}},
		},
		bson.D{
			{"$group",bson.D{
				{"_id", "$age"},
				{"name",bson.D{{"$push","$name"}}},
				{"totalStudents",bson.D{{"$sum",1}}},
			}},
		},
		bson.D{
			{"$out","totalStudents"},
		},
	}
	cursor, err := studentCollection.Aggregate(context.TODO(), pipeLine)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var students []bson.M
	if err = cursor.All(context.TODO(), &students); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func facet(ctx *gin.Context) {
	pipeLine := mongo.Pipeline{
		bson.D{{"$facet", bson.M{
			// First facet: Filtering students above age 20
			"filteredResults": bson.A{
				bson.D{{"$match", bson.D{{"age", bson.D{{"$gt", 20}}}}}},
			},
			// Second facet: Grouping students by age
			"groupedByAge": bson.A{
				bson.D{{"$group", bson.D{
					{"_id", "$age"},
					{"names", bson.D{{"$push", "$name"}}},
					{"totalStudents", bson.D{{"$sum", 1}}},
				}}},
			},
		}}},
	}

	cursor, err := studentCollection.Aggregate(context.TODO(), pipeLine)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var students []bson.M
	if err = cursor.All(context.TODO(), &students); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func main() {
	router := gin.Default()

	router.POST("/students", createStudent)
	router.POST("/students/addFields", addFields)
	router.GET("/students/count", count)
	router.GET("/students/group", group)
	router.GET("/students/match", match)
	router.GET("/students/sort", sort)
	router.GET("/students/lookup", lookup)
	router.GET("/students/merge",mergeDoc)
	router.GET("/students/out",out)
	router.GET("/students/facet",facet)
	router.Run(":3000")
}
