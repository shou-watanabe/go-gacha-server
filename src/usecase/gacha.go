package usecase

import (
	"context"
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
	"techtrain-mission/src/utils"
)

type GachaUsecase interface {
	Draw(ctx context.Context, times int, token string) ([]*entity.Chara, error)
}

type gachaUsecase struct {
	userRepo      repository.UserRepository
	charaRepo     repository.CharaRepository
	userCharaRepo repository.UserCharaRepository
}

func NewGachaUsecase(ur repository.UserRepository, cr repository.CharaRepository, ucr repository.UserCharaRepository) GachaUsecase {
	charaUsecase := gachaUsecase{userRepo: ur, charaRepo: cr, userCharaRepo: ucr}
	return &charaUsecase
}

func (gu *gachaUsecase) Draw(ctx context.Context, times int, token string) ([]*entity.Chara, error) {
	charas, err := gu.charaRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var pArray []int
	for _, chara := range charas {
		pArray = append(pArray, chara.Probability)
	}

	var gotCharas []*entity.Chara
	for i := 0; i < times; i++ {
		id := utils.WeightPick(pArray)
		gotCharas = append(gotCharas, charas[id])
	}

	ue, err := gu.userRepo.Get(ctx, token)
	if err != nil {
		return nil, err
	}

	if err := gu.userCharaRepo.Store(ctx, *ue, gotCharas); err != nil {
		return nil, err
	}

	return gotCharas, nil
}
