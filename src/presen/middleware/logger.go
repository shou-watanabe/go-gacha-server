package middleware

import (
	"net/http"

	"techtrain-mission/src/core/logger"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		logger.HttpLogging("incoming request", r)

		next.ServeHTTP(rw, r)
	})
}
