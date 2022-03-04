package middlewares

import (
	"api/src/authentication"
	"api/src/messages"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.URL, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.TokenValid(r); err != nil {
			messages.GenericError(err, "Error, invalid token")
			return
		}
		next(w, r)
	}
}
