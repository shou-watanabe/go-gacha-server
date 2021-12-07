//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	infra "techtrain-mission/src/infra/user"
	"techtrain-mission/src/presen/handler"
	"techtrain-mission/src/usecase"
)

func initUserHandler(db *gorm.DB) handler.UserHandler {
	wire.Build(
		infra.NewRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	)
	return nil
}
