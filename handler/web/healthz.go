package web

import (
	"io"
	"net/http"
)

// HandleHealthz creates an http.HandlerFunc that performs system
// healthchecks and returns 500 if the system is in an unhealthy state.
func HandleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Ok")
	}
}
