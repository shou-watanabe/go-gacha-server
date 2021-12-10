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

func initUserHandler() handler.UserHandler {
	wire.Build(
		sql.NewDB,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	)
	return nil
}
