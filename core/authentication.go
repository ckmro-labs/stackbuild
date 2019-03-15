package core

import "time"

type (
	//Authentication for 3rd auth information.
	Authentication struct {
		UserID    string    `json:"id"`
		UID       string    `json:"uid"`
		AuthName  string    `json:"name"`
		Token     string    `json:"token"`
		Expired   int64     `json:"expired"`
		CreatedAt time.Time `json:"created_at"`
		UpdateAt  time.Time `json:"updated_at"`
	}

	//AuthenticationStore store for authentication.
	AuthenticationStore interface {
	}
)
