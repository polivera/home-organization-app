package common

import "fmt"

type ErrorNotFound struct {
	Item string
}

func (e ErrorNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Item)
}

type ErrorValidation struct {
	Field string
}

func (e ErrorValidation) Error() string {
	return fmt.Sprintf("%s is not valid", e.Field)
}

type RepositoryUnexpectedError struct {
	Message string
}

func (rue RepositoryUnexpectedError) Error() string {
	return fmt.Sprintf("Unexpected repository error: %s", rue.Message)
}
