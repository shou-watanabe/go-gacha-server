package repository

import "techtrain-mission/src/domain/entity"

type UserRepository interface {
	Create(*entity.User) (*entity.User, error)
	Get(string) (*entity.User, error)
}
