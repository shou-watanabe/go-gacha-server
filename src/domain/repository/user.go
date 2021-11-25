package repository

import "techtrain-mission/src/domain/entity"

type UserRepository interface {
	Create(string) (*entity.User, error)
	Get(string) (*entity.User, error)
}
