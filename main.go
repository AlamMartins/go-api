package main

import (
	"amsolutions/config"
	"amsolutions/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// Initialize database
	db := config.InitDB(os.Getenv("MONGODB_URI"))

	// Set up routes
	router.GET("/users", handlers.GetUsers(db))
	router.POST("/users", handlers.CreateUser(db))
	router.GET("/users/:id", handlers.GetUser(db))
	router.POST("/users/login", handlers.LoginUser(db))
	router.PUT("/users/:id", handlers.UpdateUser(db))    // PUT para atualizar
	router.DELETE("/users/:id", handlers.DeleteUser(db)) // DELETE para deletar

	router.Run(":8080")
}
