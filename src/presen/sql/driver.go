package sql

import "context"

type Driver interface {
	ExecContext(context.Context, string, ...interface{}) (Result, error)
	QueryContext(context.Context, string, ...interface{}) (Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) (Row, error)
}
