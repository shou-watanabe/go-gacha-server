package main

import (
	"net/http"
	"os"
	"techtrain-mission/src/presen/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	userHandler := initUserHandler()

	handler.InitRouting(userHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
