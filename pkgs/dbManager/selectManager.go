package dbManager

import (
	"database/sql"
	"log"
)

// RunSelectProto ... todo description
func RunSelectProto(db *sql.DB, inputQuery string) map[int]string {
	var resultSet = map[int]string{}

	rows, err := db.Query(inputQuery)

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var idPortinformer int
		var portinformerDescription string

		if err := rows.Scan(&idPortinformer, &portinformerDescription); err != nil {
			log.Fatal(err)
		}

		resultSet[idPortinformer] = portinformerDescription
	}

	return resultSet
}

// runSelect ... todo description
func runSelect(db *sql.DB, inputQuery string) *sql.Rows {
	rows, err := db.Query(inputQuery)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

// GetAllPortinformers ... todo description
func GetAllPortinformers(db *sql.DB) map[int]string {
	var resultSet = map[int]string{}

	inputQuery := "SELECT id_portinformer, portinformer_description FROM portinformers"

	rows := runSelect(db, inputQuery)
	defer rows.Close()

	for rows.Next() {
		var idPortinformer int
		var portinformerDescription string

		if err := rows.Scan(&idPortinformer, &portinformerDescription); err != nil {
			log.Fatal(err)
		}
		resultSet[idPortinformer] = portinformerDescription
	}

	return resultSet
}
