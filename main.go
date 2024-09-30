package main

import (
	"go-chatgpt/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the port from environment variables
	port := os.Getenv("PORT")

	r := gin.Default()

	// Group the user routes under /users
	apiRoutes := r.Group("/api")
	routes.ChatGPTRoutes(apiRoutes)

	r.Run(":" + port) // Start the server on the specified port
}
