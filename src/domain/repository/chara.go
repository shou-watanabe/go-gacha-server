package repository

import (
	"context"

	"techtrain-mission/src/domain/entity"
)

type CharaRepository interface {
	List(ctx context.Context) ([]*entity.Chara, error)
}
