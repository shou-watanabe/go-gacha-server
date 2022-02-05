//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"

	"go-gacha-server/src/infra/repository"
	"go-gacha-server/src/presen/handler"
	"go-gacha-server/src/usecase"
)

func initUserHandler(driver *sql.DB) handler.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	)
	return nil
}

func initCharaHandler(driver *sql.DB) handler.CharaHandler {
	wire.Build(
		repository.NewUserCharaRepository,
		repository.NewUserRepository,
		usecase.NewCharaUsecase,
		handler.NewCharaHandler,
	)
	return nil
}

func initGachaHandler(driver *sql.DB) handler.GachaHandler {
	wire.Build(
		repository.NewUserRepository,
		repository.NewCharaRepository,
		repository.NewUserCharaRepository,
		usecase.NewGachaUsecase,
		handler.NewGachaHandler,
	)
	return nil
}
