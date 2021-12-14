package repository

import (
	"context"
	"log"

	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"
	infra "techtrain-mission/src/infra/sql"
	"techtrain-mission/src/presen/sql"
)

type userRepository struct {
	db sql.Driver
}

func NewUserRepository(db infra.DriverImpl) repository.UserRepository {
	return &userRepository{db: &db}
}

func (ur *userRepository) Create(ctx context.Context, name string) (*entity.User, error) {
	const (
		insert  = `INSERT INTO users (name, token) VALUE (?, UUID())`
		confirm = `SELECT * FROM users WHERE id = ?`
	)

	// stmt, err := ur.db.PrepareContext(ctx, insert)
	// if err != nil {
	// 	return nil, err
	// }

	// defer stmt.Close()

	res, err := ur.db.ExecContext(ctx, insert, name)
	// res, err := stmt.ExecContext(ctx, name)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	ue := &entity.User{}

	// stmt, err = ur.db.PrepareContext(ctx, confirm)
	// if err != nil {
	// 	return nil, err
	// }

	row, err := ur.db.QueryRowContext(ctx, confirm, id)
	if err != nil {
		return nil, err
	}

	if err = row.Scan(&ue.Id, &ue.Name, &ue.Token); err != nil {
		return nil, err
	}

	// err = stmt.QueryRowContext(ctx, id).Scan(&ue.Id, &ue.Name, &ue.Token)

	// gormバージョン
	// ur.Conn.Table("users").Exec("insert users (name, token) value (?, UUID())", name).Scan(&ue)
	// if ue.Name == "" {
	// 	err := errors.New("use not found")
	// 	return ue, err
	// }

	return ue, nil
}

func (ur *userRepository) Get(ctx context.Context, token string) (*entity.User, error) {
	const read = `SELECT * FROM users WHERE token = ?`

	// stmt, err := ur.db.PrepareContext(ctx, read)
	// if err != nil {
	// 	return nil, err
	// }

	// defer stmt.Close()

	// row := stmt.QueryRowContext(ctx, token)
	row, err := ur.db.QueryRowContext(ctx, read, token)
	if err != nil {
		log.Println("ここでエラー")
		return nil, err
	}

	ue := &entity.User{}
	err = row.Scan(&ue.Id, &ue.Name, &ue.Token)
	if err != nil {
		log.Println("こっちでエラー")
		return nil, err
	}

	// gormバージョン
	// if err := ur.Conn.Table("users").First(&user, "token = ?", token).Error; err != nil {
	// 	return nil, err
	// }

	return ue, nil
}

func (ur *userRepository) Update(ctx context.Context, name string, token string) (*entity.User, error) {
	const update = `UPDATE users Set name = ? WHERE token = ?`

	ue := &entity.User{}
	// stmt, err := ur.db.PrepareContext(ctx, update)
	// if err != nil {
	// 	return nil, err
	// }

	// defer stmt.Close()

	// _, err = stmt.ExecContext(ctx, name, token)
	_, err := ur.db.ExecContext(ctx, update, name, token)
	if err != nil {
		return nil, err
	}

	// gormバージョン
	// if err := ur.Conn.Table("users").Update("name", name).Error; err != nil {
	// 	return nil, err
	// }
	return ue, nil
}
