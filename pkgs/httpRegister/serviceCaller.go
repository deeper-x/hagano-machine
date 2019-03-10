package httpRegister

import "net/http"

// ServiceCaller .. description todo
func ServiceCaller() {
	http.HandleFunc("/regFoo", RegFoo)
	http.HandleFunc("/regBar", RegBar)
}
