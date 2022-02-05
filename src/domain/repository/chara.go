package repository

import (
	"context"

	"go-gacha-server/src/domain/entity"
)

type CharaRepository interface {
	List(ctx context.Context) ([]*entity.Chara, error)
}
