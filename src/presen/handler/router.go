package handler

import (
	"net/http"
	"techtrain-mission/src/presen/middleware"
)

func InitRouting(userHandler UserHandler) {
	create := http.HandlerFunc(userHandler.Create)
	http.Handle("/user/create", middleware.Logger(create))

	get := http.HandlerFunc(userHandler.Get)
	http.Handle("/user/get", middleware.Logger(get))

	update := http.HandlerFunc(userHandler.Update)
	http.Handle("/user/update", middleware.Logger(update))
}
