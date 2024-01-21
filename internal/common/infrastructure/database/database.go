package database

import (
	"database/sql"
)

type Connection interface {
	Open() error
	Close() error
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	QueryRow(sql string, args ...interface{}) *sql.Row
}
