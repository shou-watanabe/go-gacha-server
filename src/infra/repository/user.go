package repository

import (
	"database/sql"
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(name string) (*entity.User, error) {
	ue := &entity.User{}
	// ur.Conn.Table("users").Exec("insert users (name, token) value (?, UUID())", name).Scan(&ue)
	// if ue.Name == "" {
	// 	err := errors.New("use not found")
	// 	return ue, err
	// }
	return ue, nil
}

func (ur *userRepository) Get(token string) (*entity.User, error) {
	user := &entity.User{}
	// if err := ur.Conn.Table("users").First(&user, "token = ?", token).Error; err != nil {
	// 	return nil, err
	// }

	return user, nil
}

func (ur *userRepository) Update(name string, token string) (*entity.User, error) {
	user := &entity.User{Token: token}
	// if err := ur.Conn.Table("users").Update("name", name).Error; err != nil {
	// 	return nil, err
	// }
	return user, nil
}
