package repository

import (
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/domain/repository"

	"github.com/jinzhu/gorm"
)

type userCharaRepository struct {
	Conn *gorm.DB
}

func NewUserCharaRepository(conn *gorm.DB) repository.UserCharaRepository {
	return &userCharaRepository{Conn: conn}
}

func (ucr *userCharaRepository) List(ue entity.User) ([]*entity.UserChara, error) {
	// rows, err := ucr.Conn.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	// if err != nil {
	// 	return nil, err
	// }
	// for rows.Next() {
	// 	//...
	// }
	// 可能なら条件に合うcharaIdを取得するたびにCharaを参照する
	// User.Idからuser_character_possessionsテーブルにアクセスして、userが保持しているcharaIdを取得する
	// charaIdからcharactersテーブルにアクセスしてキャラ情報をとる
	// uce := &entity.UserChara{}
	// ur.Conn.Table("")
	return nil, nil
}
