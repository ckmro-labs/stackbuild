package core

import "net/http"

// Session management authenticated users.
type Session interface {
	// Create creates a new user session and writes the
	// session to the http.Response.
	Create(http.ResponseWriter, *User) error
	// Delete deletes the user session from the http.Response.
	Delete(http.ResponseWriter) error
	// Get returns the session from the http.Request.
	Get(*http.Request) (*User, error)
}
