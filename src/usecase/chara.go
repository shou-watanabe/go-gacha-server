package usecase

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type CharaUsecase interface {
	List(token string) ([]*entity.Chara, error)
}

type charaUsecase struct {
	charaRepo repository.CharaRepository
}

func NewCharaUsecase(cr repository.CharaRepository) CharaUsecase {
	charaUsecase := charaUsecase{charaRepo: cr}
	return &charaUsecase
}

func (cu *charaUsecase) List(token string) (chara []*entity.Chara, err error) {
	chara, err = cu.charaRepo.List(token)
	return
}
