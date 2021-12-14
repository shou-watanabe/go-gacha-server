package sql

import (
	"context"
	"database/sql"
	"log"

	presen "techtrain-mission/src/presen/sql"
)

type DriverImpl struct {
	db *sql.DB
}

func NewDriver() *DriverImpl {
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

	return &DriverImpl{db: db}
}

func (sd *DriverImpl) QueryRowContext(ctx context.Context, sql string, args ...interface{}) (presen.Row, error) {
	stmt, err := sd.db.PrepareContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, args)
	if err = row.Err(); err != nil {
		return nil, err
	}

	return &rowImpl{
		Result: row,
	}, nil
}

func (sd *DriverImpl) QueryContext(ctx context.Context, sql string, args ...interface{}) (presen.Rows, error) {
	stmt, err := sd.db.PrepareContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args)
	if err != nil {
		return nil, err
	}

	return &rowsImpl{
		Result: rows,
	}, nil
}

func (sd *DriverImpl) ExecContext(ctx context.Context, sql string, args ...interface{}) (presen.Result, error) {
	stmt, err := sd.db.PrepareContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args)
	if err != nil {
		return nil, err
	}

	return res, nil
}
