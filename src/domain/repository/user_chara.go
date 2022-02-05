package repository

import (
	"context"

	"go-gacha-server/src/domain/entity"
)

type UserCharaRepository interface {
	List(ctx context.Context, ue entity.User) ([]*entity.UserChara, error)
	Store(ctx context.Context, ue entity.User, ces []*entity.Chara) error
}
