package repository

import (
	"context"
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

func (ur *userRepository) Create(ctx context.Context, name string) (*entity.User, error) {
	const (
		insert  = `INSERT INTO users (name, token) VALUE (?, UUID())`
		confirm = `SELECT * FROM users WHERE id = ?`
	)

	stmt, err := ur.db.PrepareContext(ctx, insert)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, name)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	ue := &entity.User{}

	stmt, err = ur.db.PrepareContext(ctx, confirm)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&ue.Id, &ue.Name, &ue.Token)
	if err != nil {
		return nil, err
	}

	return ue, nil
}

func (ur *userRepository) Get(ctx context.Context, token string) (*entity.User, error) {
	user := &entity.User{}
	// if err := ur.Conn.Table("users").First(&user, "token = ?", token).Error; err != nil {
	// 	return nil, err
	// }

	return user, nil
}

func (ur *userRepository) Update(ctx context.Context, name string, token string) (*entity.User, error) {
	user := &entity.User{Token: token}
	// if err := ur.Conn.Table("users").Update("name", name).Error; err != nil {
	// 	return nil, err
	// }
	return user, nil
}
