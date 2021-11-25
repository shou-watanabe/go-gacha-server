package usecase

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type UserUsecase interface {
	Create(string) (*entity.User, error)
	Get(string) (*entity.User, error)
	Update(string, string) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo: ur}
	return &userUsecase
}

func (uu *userUsecase) Create(name string) (user *entity.User, err error) {
	user, err = uu.userRepo.Create(name)
	return
}

func (uu *userUsecase) Get(token string) (user *entity.User, err error) {
	user, err = uu.userRepo.Get(token)
	return
}

func (uu *userUsecase) Update(name string, token string) (user *entity.User, err error) {
	user, err = uu.userRepo.Update(name, token)
	return
}
