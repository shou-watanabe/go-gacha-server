package usecase

import (
	"context"
	"go-gacha-server/src/domain/entity"
	"go-gacha-server/src/domain/repository"
)

type CharaUsecase interface {
	List(ctx context.Context, token string) ([]*entity.UserChara, error)
}

type charaUsecase struct {
	userRepo      repository.UserRepository
	userCharaRepo repository.UserCharaRepository
}

func NewCharaUsecase(ur repository.UserRepository, ucr repository.UserCharaRepository) CharaUsecase {
	charaUsecase := charaUsecase{userRepo: ur, userCharaRepo: ucr}
	return &charaUsecase
}

func (cu *charaUsecase) List(ctx context.Context, token string) ([]*entity.UserChara, error) {
	user, err := cu.userRepo.Get(ctx, token)
	if err != nil {
		return nil, err
	}

	userChara, err := cu.userCharaRepo.List(ctx, *user)
	if err != nil {
		return nil, err
	}
	return userChara, nil
}
