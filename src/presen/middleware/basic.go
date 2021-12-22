package middleware

import (
	"net/http"
	"os"
)

func Basic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, pass, ok := r.BasicAuth()
		uid := os.Getenv("BASIC_AUTH_USER_ID")
		up := os.Getenv("BASIC_AUTH_PASSWORD")
		if uid != name || up != pass || !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
