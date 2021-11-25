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

func (ur *userRepository) Get(token string) (*entity.User, error) {
	user := &entity.User{}
	if err := ur.Conn.Table("users").First(&user, "token = ?", token).Error; err != nil {
		return nil, err
	}

	return user, nil
}
