package sql

import (
	"context"
	"database/sql"
	"log"
)

type SqlDriver struct {
	Db *sql.DB
}

func NewDriver() *SqlDriver {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		log.Println("DB connect failed")
		panic(err)
	}

	log.Println("DB connect success")

	return &SqlDriver{Db: db}
}

func QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row {
	return nil
}

func QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return nil, nil
}
