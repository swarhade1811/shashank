package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/data")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createTable() {
	createQuery := `CREATE TABLE IF NOT EXISTS feed_configurations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		location VARCHAR(255) NOT NULL,
		index_name VARCHAR(255) NOT NULL,
		frequency VARCHAR(255) NOT NULL,
		columns VARCHAR(255) NOT NULL,
		cap INT NOT NULL,
		scrub BOOLEAN NOT NULL,
		look_a_like_config VARCHAR(255) NOT NULL
	);`

	_, err := db.Exec(createQuery)
	if err != nil {
		log.Fatal(err)
	}
}
