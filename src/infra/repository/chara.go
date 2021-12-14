package repository

import (
	"database/sql"
	"techtrain-mission/src/domain/repository"
)

type charaRepository struct {
	db *sql.DB
}

func NewCharaRepository(db *sql.DB) repository.CharaRepository {
	return &charaRepository{db: db}
}

// func (cr *charaRepository) List(token string) ([]*entity.Chara, error) {
// 	return nil, nil
// }
