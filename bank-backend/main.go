package main

import (
	"bank-backend/api"
	"bank-backend/middleware"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize Gin (production mode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORS())

	// Routes
	api.RegisterRoutes(r)

	// Start server
	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
