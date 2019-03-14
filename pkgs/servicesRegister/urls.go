package servicesRegister

import "net/http"

// UrlsRegister .. description todo
func UrlsRegister() {
	http.HandleFunc("/regFoo", RegFoo)
	http.HandleFunc("/regBar", RegBar)
	http.HandleFunc("/AllPortinformers", AllPortinformers)
	http.HandleFunc("/AllPlannedArrivals", AllPlannedArrivals)
}
