package usecase

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type CharaUsecase interface {
	Get(id int) (*entity.Chara, error)
}

type charaUsecase struct {
	charaRepo repository.CharaRepository
}

func NewCharaUsecase(cr repository.CharaRepository) CharaUsecase {
	charaUsecase := charaUsecase{charaRepo: cr}
	return &charaUsecase
}

func (uu *charaUsecase) Get(id int) (chara *entity.Chara, err error) {
	chara, err = uu.charaRepo.Get(id)
	return
}
