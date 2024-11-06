package middleware

import (
	"net/http"
)

func Role(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// logic role
		next.ServeHTTP(w, r)
	})
}
