package usecase

import (
	"context"
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
	"techtrain-mission/src/utils"
)

type GachaUsecase interface {
	Draw(ctx context.Context, times int) ([]*entity.Chara, error)
}

type gachaUsecase struct {
	charaRepo repository.CharaRepository
}

func NewGachaUsecase(cr repository.CharaRepository) GachaUsecase {
	charaUsecase := gachaUsecase{charaRepo: cr}
	return &charaUsecase
}

func (gu *gachaUsecase) Draw(ctx context.Context, times int) ([]*entity.Chara, error) {
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

	return gotCharas, nil
}
