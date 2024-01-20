package database

import (
	"database/sql"
)

type Connection interface {
	Open() error
	Connect() (*sql.Conn, error)
	Query(sql string, args ...any) (*sql.Rows, error)
}
