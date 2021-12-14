package main

import (
	"net/http"
	"os"
	"techtrain-mission/src/infra/sql"
	"techtrain-mission/src/presen/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := sql.NewDriver()
	userHandler := initUserHandler(*db)
	charaHandler := initCharaHandler(*db)

	handler.InitRouting(userHandler, charaHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
