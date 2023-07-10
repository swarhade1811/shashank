package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	// Database connection setup
	db = setupDatabase()
	defer db.Close()

	// Create the "Feed configuration" table if it doesn't exist
	createTable()

	// Initialize the router
	router := mux.NewRouter()

	// corsMiddleware := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders:   []string{"Content-Type"},
	// 	AllowCredentials: true,
	// })
	corsMiddleware := cors.Default()

	router.Use(corsMiddleware.Handler)

	// Define API routes
	router.HandleFunc("/feed-configurations", getAllFeedConfigurations).Methods("GET")
	router.HandleFunc("/feed-configurations/{id}", getFeedConfiguration).Methods("GET")
	router.HandleFunc("/feed-configurations", createFeedConfiguration).Methods("POST")
	router.HandleFunc("/feed-configurations/{id}", updateFeedConfiguration).Methods("PUT")
	router.HandleFunc("/feed-configurations/{id}", deleteFeedConfiguration).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
