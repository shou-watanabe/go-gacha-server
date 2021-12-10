package repository

import "techtrain-mission/src/domain/entity"

type CharaRepository interface {
	Get(id int) (*entity.Chara, error)
}
