package usecase

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type CharaUsecase interface {
	List(token string) ([]*entity.UserChara, error)
}

type charaUsecase struct {
	userRepo      repository.UserRepository
	userCharaRepo repository.UserCharaRepository
}

func NewCharaUsecase(ur repository.UserRepository, ucr repository.UserCharaRepository) CharaUsecase {
	charaUsecase := charaUsecase{userRepo: ur, userCharaRepo: ucr}
	return &charaUsecase
}

func (cu *charaUsecase) List(token string) ([]*entity.UserChara, error) {
	user, err := cu.userRepo.Get(token)
	if err != nil {
		return nil, err
	}

	userChara, err := cu.userCharaRepo.List(*user)
	if err != nil {
		return nil, err
	}
	return userChara, nil
}
