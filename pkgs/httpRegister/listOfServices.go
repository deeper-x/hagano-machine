package httpRegister

import (
	"fmt"
	"html"
	"net/http"
)

// RegFoo ... todo description
var RegFoo = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here I am... %q", html.EscapeString(r.URL.Path))
}

// RegBar ... todo description
var RegBar = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here I am ... %q", html.EscapeString(r.URL.Path))
}
