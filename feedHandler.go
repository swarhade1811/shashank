package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FeedConfiguration struct {
	ID              int    `json:"id"`
	Location        string `json:"location"`
	IndexName       string `json:"index_name"`
	Frequency       string `json:"frequency"`
	Columns         string `json:"columns"`
	Cap             int    `json:"cap"`
	Scrub           bool   `json:"scrub"`
	LookALikeConfig string `json:"look_a_like_config"`
}

// Handler to get all feed configurations
func getAllFeedConfigurations(w http.ResponseWriter, r *http.Request) {
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
	var fc FeedConfiguration
	_ = json.NewDecoder(r.Body).Decode(&fc)

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
