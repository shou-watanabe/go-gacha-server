package repository

import (
	"errors"
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

func (ur *userRepository) Create(name string) (ue *entity.User, err error) {
	ur.Conn.Table("users").Exec("insert users (name, token) value (?, UUID())", name).Scan(&ue)
	if ue.Name == "" {
		err = errors.New("use not found")
	}
	return
}

func (ur *userRepository) Get(token string) (*entity.User, error) {
	user := &entity.User{}
	if err := ur.Conn.Table("users").First(&user, "token = ?", token).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) Update(name string, token string) (*entity.User, error) {
	user := &entity.User{Token: token}
	if err := ur.Conn.Table("users").Update("name", name).Error; err != nil {
		return nil, err
	}
	return user, nil
}
