package usecase

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type UserUsecase interface {
	Create(string) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo: ur}
	return &userUsecase
}

func (usecase *userUsecase) Create(name string) (user *entity.User, err error) {
	ue := &entity.User{Name: name}
	user, err = usecase.userRepo.Create(ue)
	return
}
