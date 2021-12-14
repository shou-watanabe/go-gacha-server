package repository

import (
	"context"

	"techtrain-mission/src/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, name string) (*entity.User, error)
	Get(ctx context.Context, token string) (*entity.User, error)
	Update(ctx context.Context, name string, token string) (*entity.User, error)
}
