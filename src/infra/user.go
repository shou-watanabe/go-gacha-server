package infra

import (
	"errors"
	userE "techtrain-mission/src/domain/entity/user"
	userR "techtrain-mission/src/domain/repository/user"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) userR.Repository {
	return &userRepository{Conn: conn}
}

func (ur *userRepository) Create(name string) (*userE.Entity, error) {
	ue := &userE.Entity{}
	ur.Conn.Table("users").Exec("insert users (name, token) value (?, UUID())", name).Scan(&ue)
	if ue.Name == "" {
		err := errors.New("use not found")
		return ue, err
	}
	return ue, nil
}

func (ur *userRepository) Get(token string) (*userE.Entity, error) {
	user := &userE.Entity{}
	if err := ur.Conn.Table("users").First(&user, "token = ?", token).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) Update(name string, token string) (*userE.Entity, error) {
	user := &userE.Entity{Token: token}
	if err := ur.Conn.Table("users").Update("name", name).Error; err != nil {
		return nil, err
	}
	return user, nil
}
