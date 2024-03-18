package http

import "net/http"

// RespondInJson sets the 'Content-Type' as 'application/json' in the headers.
func RespondInJson(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}
