package handler

import (
	"net/http"
	"techtrain-mission/src/presen/middleware"
)

func InitRouting(userHandler UserHandler, charaHandler CharaHandler) {
	userCreate := http.HandlerFunc(userHandler.Create)
	http.Handle("/user/create", middleware.Logger(userCreate))

	userGet := http.HandlerFunc(userHandler.Get)
	http.Handle("/user/get", middleware.Logger(userGet))

	userUpdate := http.HandlerFunc(userHandler.Update)
	http.Handle("/user/update", middleware.Logger(userUpdate))

	characterList := http.HandlerFunc(charaHandler.List)
	http.Handle("/character/list", middleware.Logger(characterList))
}
