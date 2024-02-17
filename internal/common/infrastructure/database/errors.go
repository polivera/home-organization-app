package database

import "fmt"

type ErrorForeignRowNotFound struct {
	Table string
}

func (e ErrorForeignRowNotFound) Error() string {
	return fmt.Sprintf("foreign row on table `%s` not found", e.Table)
}

var ErrForeignRowNotFound = &ErrorForeignRowNotFound{}

type ErrorDuplicateKey struct {
	Key string
}

func (e ErrorDuplicateKey) Error() string {
	return fmt.Sprintf("record already exist under key `%s`", e.Key)
}

var ErrDuplicateKey = &ErrorDuplicateKey{}
