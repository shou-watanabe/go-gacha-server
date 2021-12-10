package repository

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"

	"github.com/jinzhu/gorm"
)

type userCharaRepository struct {
	Conn *gorm.DB
}

func NewUserCharaRepository(conn *gorm.DB) repository.UserCharaRepository {
	return &userCharaRepository{Conn: conn}
}

func (ucr *userCharaRepository) List(token string) ([]*entity.UserChara, error) {
	return nil, nil
}
