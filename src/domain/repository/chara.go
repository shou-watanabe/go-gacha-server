package repository

import "techtrain-mission/src/domain/entity"

type CharaRepository interface {
	Get(int) (*entity.Chara, error)
}
