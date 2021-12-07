package usecase

import (
	userE "techtrain-mission/src/domain/entity/user"
	userR "techtrain-mission/src/domain/repository/user"
)

type UserUsecase interface {
	Create(string) (*userE.Entity, error)
	Get(string) (*userE.Entity, error)
	Update(string, string) (*userE.Entity, error)
}

type userUsecase struct {
	userRepo userR.Repository
}

func NewUserUsecase(ur userR.Repository) UserUsecase {
	userUsecase := userUsecase{userRepo: ur}
	return &userUsecase
}

func (uu *userUsecase) Create(name string) (user *userE.Entity, err error) {
	user, err = uu.userRepo.Create(name)
	return
}

func (uu *userUsecase) Get(token string) (user *userE.Entity, err error) {
	user, err = uu.userRepo.Get(token)
	return
}

func (uu *userUsecase) Update(name string, token string) (user *userE.Entity, err error) {
	user, err = uu.userRepo.Update(name, token)
	return
}
