package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

const uri = "mongodb+srv://bhosalerupesh67:L2hSiWX3IeytWQNm@jwt-users.8jhul.mongodb.net/?retryWrites=true&w=majority&appName=JWT-Users"

var mongoClient *mongo.Client
var userCollection *mongo.Collection

func init() {
	if err := connect_to_mongo(); err != nil {
		fmt.Println("could not connect to mongo")
	} else {
		fmt.Println("connected to mongo")
	}
	userCollection = mongoClient.Database("users").Collection("users")
}

func connect_to_mongo() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}

var secretKey = []byte("secretKey")

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(secretKey)
}

func main() {
	router := gin.Default()

	router.POST("/register", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingUser User
		err := userCollection.FindOne(context.TODO(), bson.M{
			"$or": []bson.M{
				{"username": user.Username},
				{"email": user.Email},
			},
		}).Decode(&existingUser)
		if err == nil {
			ctx.JSON(400, gin.H{"error": "user already exists"})
			return
		} else if err != mongo.ErrNoDocuments {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
			return
		}
		user.Password = string(hashedPassword)

		insertResult, err := userCollection.InsertOne(context.TODO(), user)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, gin.H{"message": "user created", "id": insertResult.InsertedID})
	})

	router.GET("/", func(ctx *gin.Context) {
		allUsers, err := userCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var users []User
		if err = allUsers.All(context.TODO(), &users); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"users": users})
	})

	router.POST("/login", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		var foundUser User
		err := userCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&foundUser)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid username/password"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid username/password"})
			return
		}

		token, err := generateToken(foundUser.Username)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "could not generate token"})
			return
		}
		ctx.SetCookie("token", token, 3600, "/", "", false, false)
		cookie, error := ctx.Cookie("token")
		if error != nil {
			ctx.JSON(500, gin.H{"error": "could not set cookie"})
			return
		}

		ctx.JSON(200, gin.H{"message": "login successful", "token": token, "cookie": cookie})
	})

	router.DELETE("/deleteUser/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")

		var foundUser User
		error := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&foundUser)
		if error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid username"})
			return
		}
		_, error = userCollection.DeleteOne(context.TODO(), bson.M{"username":username})
		if error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
	})

	router.Run(":3000")
}
