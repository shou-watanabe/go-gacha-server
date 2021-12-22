package middleware

import (
	"net/http"
)

func Layers(handler http.Handler) http.Handler {
	return Recovery(
		Context(
			Logger(
				handler,
			),
		),
	)
}

func AuthLayers(handler http.Handler) http.Handler {
	return Recovery(
		Context(
			Logger(
				Basic(
					handler,
				),
			),
		),
	)
}
