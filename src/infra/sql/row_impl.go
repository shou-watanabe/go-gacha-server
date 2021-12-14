package sql

import (
	"database/sql"
)

type rowImpl struct {
	Result *sql.Row
}

func (r *rowImpl) Scan(args ...interface{}) error {
	err := r.Result.Scan(args...)
	if err != nil {
		return err
	}

	return nil
}
