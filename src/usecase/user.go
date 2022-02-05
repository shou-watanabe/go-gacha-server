package usecase

import (
	"context"
	"go-gacha-server/src/domain/entity"
	"go-gacha-server/src/domain/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, name string) (*entity.User, error)
	Get(ctx context.Context, token string) (*entity.User, error)
	Update(ctx context.Context, name string, token string) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo: ur}
	return &userUsecase
}

func (uu *userUsecase) Create(ctx context.Context, name string) (user *entity.User, err error) {
	user, err = uu.userRepo.Create(ctx, name)
	return
}

func (uu *userUsecase) Get(ctx context.Context, token string) (user *entity.User, err error) {
	user, err = uu.userRepo.Get(ctx, token)
	return
}

func (uu *userUsecase) Update(ctx context.Context, name string, token string) (user *entity.User, err error) {
	user, err = uu.userRepo.Update(ctx, name, token)
	return
}
