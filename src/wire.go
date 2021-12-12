//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"techtrain-mission/src/infra/repository"
	"techtrain-mission/src/presen/handler"
	"techtrain-mission/src/usecase"
)

func initUserHandler(db *gorm.DB) handler.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	)
	return nil
}

func initCharaHandler(db *gorm.DB) handler.CharaHandler {
	wire.Build(
		repository.NewUserCharaRepository,
		repository.NewUserRepository,
		usecase.NewCharaUsecase,
		handler.NewCharaHandler,
	)
	return nil
}
