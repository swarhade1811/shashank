package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FeedConfiguration struct {
	ID              int    `json:"Id"`
	Location        string `json:"Location"`
	IndexName       string `json:"IndexName"`
	Frequency       string `json:"Frequency"`
	Columns         string `json:"Columns"`
	Cap             string `json:"Cap"`
	Scrub           string `json:"Scrub"`
	LookALikeConfig string `json:"Lookalikeconfig"`
}

// Handler to get all feed configurations
func getAllFeedConfigurations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "localhost:3000")
	feedConfigurations := []FeedConfiguration{}

	rows, err := db.Query("SELECT * FROM feed_configurations")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var fc FeedConfiguration
		err := rows.Scan(&fc.ID, &fc.Location, &fc.IndexName, &fc.Frequency, &fc.Columns, &fc.Cap, &fc.Scrub, &fc.LookALikeConfig)
		if err != nil {
			log.Fatal(err)
		}
		feedConfigurations = append(feedConfigurations, fc)
	}

	json.NewEncoder(w).Encode(feedConfigurations)
}

// Handler to get a specific feed configuration by ID
func getFeedConfiguration(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var fc FeedConfiguration

	err := db.QueryRow("SELECT * FROM feed_configurations WHERE id = ?", id).
		Scan(&fc.ID, &fc.Location, &fc.IndexName, &fc.Frequency, &fc.Columns, &fc.Cap, &fc.Scrub, &fc.LookALikeConfig)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(fc)
}

// Handler to create a new feed configuration
func createFeedConfiguration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Success")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, x-requested-with")
	w.Header().Set("Content-Type", "application/json")

	var fc FeedConfiguration
	_ = json.NewDecoder(r.Body).Decode(&fc)
	fmt.Println(fc)

	insertQuery := `INSERT INTO feed_configurations (location, index_name, frequency, columns, cap, scrub, look_a_like_config)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(insertQuery, fc.Location, fc.IndexName, fc.Frequency, fc.Columns, fc.Cap, fc.Scrub, fc.LookALikeConfig)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(fc)
}

// Handler to update a feed configuration by ID
func updateFeedConfiguration(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var fc FeedConfiguration
	_ = json.NewDecoder(r.Body).Decode(&fc)

	updateQuery := `UPDATE feed_configurations SET location = ?, index_name = ?, frequency = ?, columns = ?, cap = ?, scrub = ?, look_a_like_config = ?
		WHERE id = ?`

	_, err := db.Exec(updateQuery, fc.Location, fc.IndexName, fc.Frequency, fc.Columns, fc.Cap, fc.Scrub, fc.LookALikeConfig, id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(fc)
}

// Handler to delete a feed configuration by ID
func deleteFeedConfiguration(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM feed_configurations WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Feed configuration deleted successfully")
}
