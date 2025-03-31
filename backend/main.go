package main

import (
	"bank-app/pkg/api"
	"bank-app/pkg/database"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	db, err := database.NewMySQLConnection("user:password@tcp(mysql:3306)/bankdb")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup API
	router := api.NewRouter(db)

	// Start server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
