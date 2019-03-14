package dbManager

import (
	"database/sql"
	"fmt"
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

// GetPlannedArrivals ... todo description
func GetPlannedArrivals(db *sql.DB, idPortinformer string) map[int]string {
	var resultSet = map[int]string{}

	inputQuery := fmt.Sprintf("SELECT id_planned_arrival, ts_arrival_prevision FROM planned_arrivals WHERE fk_portinformer = %s", idPortinformer)

	rows := runSelect(db, inputQuery)

	defer rows.Close()

	for rows.Next() {
		var idPlannedArrival int
		var tsArrivalPrevision string

		if err := rows.Scan(&idPlannedArrival, &tsArrivalPrevision); err != nil {
			log.Fatal(err)
		}

		resultSet[idPlannedArrival] = tsArrivalPrevision
	}

	return resultSet
}
