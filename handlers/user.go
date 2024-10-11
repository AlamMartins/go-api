package handlers

import (
	"amsolutions/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := db.Database("reactgram").Collection("users")
		var users []models.User
		cursor, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(context.TODO())
		if err = cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(users) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No users found"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetUser(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		collection := db.Database("reactgram").Collection("users")
		var user models.User
		err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		collection := db.Database("reactgram").Collection("users")
		_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		collection := db.Database("reactgram").Collection("users")
		_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

func LoginUser(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		collection := db.Database("reactgram").Collection("users")
		err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(db *mongo.Client) gin.HandlerFunc {

	return func(c *gin.Context) {

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := db.Database("test").Collection("users")
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	}

}
