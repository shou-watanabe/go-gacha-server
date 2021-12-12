package repository

import "techtrain-mission/src/domain/entity"

type UserCharaRepository interface {
	List(ue entity.User) ([]*entity.UserChara, error)
}
