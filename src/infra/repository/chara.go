package repository

import (
	"techtrain-mission/src/domain/repository"

	"github.com/jinzhu/gorm"
)

type charaRepository struct {
	Conn *gorm.DB
}

func NewCharaRepository(conn *gorm.DB) repository.CharaRepository {
	return &charaRepository{Conn: conn}
}

// func (cr *charaRepository) List(token string) ([]*entity.Chara, error) {
// 	return nil, nil
// }
