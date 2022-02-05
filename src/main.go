package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"go-gacha-server/src/infra/sql"
	"go-gacha-server/src/presen/handler"
	"go-gacha-server/src/utils"
)

func main() {
	db := sql.NewDriver()
	userHandler := initUserHandler(db)
	charaHandler := initCharaHandler(db)
	gachaHandler := initGachaHandler(db)

	handler.InitRouting(userHandler, charaHandler, gachaHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr: ":" + port,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("Server closed with error:", err)
		}
	}()

	utils.WaitSignal()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown:", err)
	}
	log.Println("Server shutdown")
}
