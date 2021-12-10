package logger

import (
	"net/http"

	"go.uber.org/zap"
)

func HttpLogging(msg string, r *http.Request) {
	logger, _ := zap.NewProduction()
	logger.Info(
		msg,
		zap.String("method", r.Method),
		zap.String("host", r.Host),
		zap.String("path", r.URL.Path),
	)
}
