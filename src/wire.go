//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"

	"techtrain-mission/src/infra/repository"
	"techtrain-mission/src/presen/handler"
	"techtrain-mission/src/usecase"
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
		repository.NewCharaRepository,
		usecase.NewCharaUsecase,
		handler.NewCharaHandler,
	)
	return nil
}
