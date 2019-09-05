package routes

import (
	"net/http"
)

// SetCORSHeaders sets headers required by CORS ajax request
func SetCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "86400")
}
