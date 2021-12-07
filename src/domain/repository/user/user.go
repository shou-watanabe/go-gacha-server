package repository

import "techtrain-mission/src/domain/entity/user"

type Repository interface {
	Create(string) (*user.Entity, error)
	Get(string) (*user.Entity, error)
	Update(string, string) (*user.Entity, error)
}
