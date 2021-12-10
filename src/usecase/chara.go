package usecase

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type CharaUsecase interface {
	List(token string) ([]*entity.UserChara, error)
}

type charaUsecase struct {
	charaRepo     repository.CharaRepository
	userCharaRepo repository.UserCharaRepository
}

func NewCharaUsecase(cr repository.CharaRepository, ucr repository.UserCharaRepository) CharaUsecase {
	charaUsecase := charaUsecase{charaRepo: cr, userCharaRepo: ucr}
	return &charaUsecase
}

func (cu *charaUsecase) List(token string) (userChara []*entity.UserChara, err error) {
	userChara, err = cu.userCharaRepo.List(token)
	return
}
