package database

import (
	"context"
	"database/sql"
)

type Connection interface {
	Open() error
	GetDB() *sql.DB
	Connect(ctx context.Context) (*sql.Conn, error)
}
