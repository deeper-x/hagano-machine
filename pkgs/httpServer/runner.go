package httpServer

import (
	"log"
	"net/http"
)

// Run ... todo description
func Run(portParameter string) {
	log.Fatal(http.ListenAndServe(portParameter, nil))
}
