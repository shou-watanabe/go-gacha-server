package logger

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	switch os.Getenv("TEST_ENV") {
	case "production", "staging":
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func HttpLogging(r *http.Request) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"host":   r.Host,
		"path":   r.URL.Path,
		"Ua":     r.UserAgent,
	})
}
