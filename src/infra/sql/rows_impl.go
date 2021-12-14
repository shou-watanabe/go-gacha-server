package sql

import (
	"database/sql"
)

type rowsImpl struct {
	Result *sql.Rows
}

func (r *rowsImpl) Close() {
	if err := r.Result.Close(); err != nil {
		panic(err)
	}
}

func (r *rowsImpl) Next() bool {
	return r.Result.Next()
}

func (r *rowsImpl) Scan(args ...interface{}) error {
	if err := r.Result.Scan(args...); err != nil {
		return err
	}

	return nil
}
