//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package database

import (
	"database/sql"
)

type Connection interface {
	Open() error
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Execute(query string, args ...interface{}) (sql.Result, error)
}
