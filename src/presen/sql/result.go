package sql

type Result interface {
	LastInsertId() (int64, error)
}
