package sql

import "database/sql"

type ResultImpl struct {
	Result sql.Result
}

func (r ResultImpl) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}
