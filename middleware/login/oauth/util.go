package oauth

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// default cookie name.
const cookieName = "_oauth_state_"

func createState(w http.ResponseWriter) string {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  random(),
		MaxAge: 1800,
	}
	http.SetCookie(w, cookie)
	return cookie.Value
}

func validateState(r *http.Request, state string) error {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return err
	}
	if state != cookie.Value {
		return ErrState
	}
	return nil
}

func deleteState(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    cookieName,
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})
}

func random() string {
	return fmt.Sprintf("%x", rand.Uint64())
}
