package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	// Define API routes
	router.HandleFunc("/feed-configurations", getAllFeedConfigurations).Methods("GET")
	router.HandleFunc("/feed-configurations/{id}", getFeedConfiguration).Methods("GET")
	router.HandleFunc("/feed-configurations", createFeedConfiguration).Methods("POST")
	router.HandleFunc("/feed-configurations/{id}", updateFeedConfiguration).Methods("PUT")
	router.HandleFunc("/feed-configurations/{id}", deleteFeedConfiguration).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
