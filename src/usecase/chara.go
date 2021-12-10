package usecase

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type CharaUsecase interface {
	List(id int) (*entity.Chara, error)
}

type charaUsecase struct {
	charaRepo repository.CharaRepository
}

func NewCharaUsecase(cr repository.CharaRepository) CharaUsecase {
	charaUsecase := charaUsecase{charaRepo: cr}
	return &charaUsecase
}

func (cu *charaUsecase) List(id int) (chara *entity.Chara, err error) {
	chara, err = cu.charaRepo.List(id)
	return
}
