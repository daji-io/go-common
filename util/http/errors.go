package http

import (
	"net/http"
	"time"
)

type ErrorResponse struct {
	TraceId  string    `json:"trace_id"`
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

// GenericBadRequest returns a standardised bad request message if no 'err' param is provided.
func GenericBadRequest(w http.ResponseWriter, err string) {
	if err == "" {
		http.Error(w, "bad request", 400)
		return
	}

	http.Error(w, err, 400)
}

func GenericInternalException(w http.ResponseWriter, r *http.Request, err string) {
	if err == "" {
		http.Error(w, "internal health error", 500)
		return
	}

	http.Error(w, err, 500)
}
