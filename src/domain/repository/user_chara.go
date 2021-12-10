package repository

import "techtrain-mission/src/domain/entity"

type UserCharaRepository interface {
	List(token string) ([]*entity.UserChara, error)
}
