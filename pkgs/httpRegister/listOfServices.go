package httpRegister

import (
	"fmt"
	"html"
	"net/http"

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

// RegFoo ... A dummy test
var RegFoo = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here I am... %q", html.EscapeString(r.URL.Path))
}

// RegBar ... Another dummy test
var RegBar = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here I am ... %q", html.EscapeString(r.URL.Path))
}
