//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"techtrain-mission/src/infra/repository"
	"techtrain-mission/src/infra/sql"
	"techtrain-mission/src/presen/handler"
	"techtrain-mission/src/usecase"
)

func initUserHandler(driver *sql.SqlDriver) handler.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	)
	return nil
}

func initCharaHandler(driver *sql.SqlDriver) handler.CharaHandler {
	wire.Build(
		repository.NewUserCharaRepository,
		repository.NewUserRepository,
		usecase.NewCharaUsecase,
		handler.NewCharaHandler,
	)
	return nil
}
