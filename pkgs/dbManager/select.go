package dbManager

import (
	"database/sql"
	"log"
)

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
