package repository

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn: conn}
}

func (ur *userRepository) Create(ue *entity.User) (*entity.User, error) {
	ur.Conn.Table("users").Exec("insert users (name, token) value (?, UUID())", ue.Name).Scan(&ue)

	return ue, nil
}
