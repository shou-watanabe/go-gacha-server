package repository

import (
	"database/sql"
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type userCharaRepository struct {
	db *sql.DB
}

func NewUserCharaRepository(db *sql.DB) repository.UserCharaRepository {
	return &userCharaRepository{db: db}
}

func (ucr *userCharaRepository) List(token string) ([]*entity.UserChara, error) {
	return nil, nil
}
