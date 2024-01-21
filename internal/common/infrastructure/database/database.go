package database

import (
	"database/sql"
)

type Connection interface {
	Open() error
	Close() error
	Query(sql string) (*sql.Rows, error)
	QueryWithParams(sql string, args ...any) (*sql.Rows, error)
}
