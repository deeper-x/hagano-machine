package dbManager

import (
	"database/sql"
	"log"
)

// GetConnection ...
func GetConnection(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// runSelect ... todo description
func runSelect(db *sql.DB, inputQuery string) *sql.Rows {
	rows, err := db.Query(inputQuery)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}
