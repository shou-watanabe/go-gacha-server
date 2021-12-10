package repository

import "techtrain-mission/src/domain/entity"

type CharaRepository interface {
	List(id int) (*entity.Chara, error)
}
