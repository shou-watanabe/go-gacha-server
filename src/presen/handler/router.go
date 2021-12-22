package handler

import (
	"net/http"
	"techtrain-mission/src/presen/middleware"
)

func InitRouting(userHandler UserHandler, charaHandler CharaHandler, gachaHandler GachaHandler) {
	userCreate := http.HandlerFunc(userHandler.Create)
	http.Handle("/user/create", middleware.AuthLayers(userCreate))

	userGet := http.HandlerFunc(userHandler.Get)
	http.Handle("/user/get", middleware.Layers(userGet))

	userUpdate := http.HandlerFunc(userHandler.Update)
	http.Handle("/user/update", middleware.AuthLayers(userUpdate))

	characterList := http.HandlerFunc(charaHandler.List)
	http.Handle("/character/list", middleware.AuthLayers(characterList))

	gachaDraw := http.HandlerFunc(gachaHandler.Draw)
	http.Handle("/gacha/draw", middleware.AuthLayers(gachaDraw))
}
