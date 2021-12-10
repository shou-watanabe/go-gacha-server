package repository

import "techtrain-mission/src/domain/entity"

type UserRepository interface {
	Create(name string) (*entity.User, error)
	Get(token string) (*entity.User, error)
	Update(name string, token string) (*entity.User, error)
}
