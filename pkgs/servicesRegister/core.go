package servicesRegister

import (
	"fmt"
	"html"
	"net/http"
	"net/url"

	"github.com/deeper-x/hagano-machine/pkgs/dbManager"
	"github.com/deeper-x/hagano-machine/pkgs/strFormatters"

	// postgres driver - used from dbManager
	_ "github.com/lib/pq"
)

var connStr = strFormatters.ConnStr
var db = dbManager.GetConnection(connStr)

// AllPortinformers ... get all portinformers list
var AllPortinformers = func(w http.ResponseWriter, r *http.Request) {
	resultSet := dbManager.GetAllPortinformers(db)
	for k, v := range resultSet {
		fmt.Fprintf(w, "<div>%d -> %v</div>", k, v)
	}
}

// AllPlannedArrivals ... get all planned arrivals
var AllPlannedArrivals = func(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idPortinformer := params.Get("id_portinformer")
	resultSet := dbManager.GetPlannedArrivals(db, idPortinformer)
	for k, v := range resultSet {
		fmt.Fprintf(w, "%d, %v", k, v)
	}
}

func parseParams(inRequest *http.Request) url.Values {
	return inRequest.URL.Query()
}

// RegFoo ... A dummy test
var RegFoo = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here I am... %q", html.EscapeString(r.URL.Path))
}

// RegBar ... Another dummy test
var RegBar = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here I am ... %q", html.EscapeString(r.URL.Path))
}
