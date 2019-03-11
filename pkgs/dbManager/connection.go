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
