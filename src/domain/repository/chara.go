package repository

import "techtrain-mission/src/domain/entity"

type CharaRepository interface {
	List(token string) ([]*entity.Chara, error)
}
