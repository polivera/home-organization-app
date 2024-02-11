package user

import "fmt"

type ErrorUserExist struct {
	Email string
}

func (e ErrorUserExist) Error() string {
	return fmt.Sprintf("User with email %s already exist", e.Email)
}

type ErrorCantHashField struct {
	Field string
}

func (e ErrorCantHashField) Error() string {
	return fmt.Sprintf("Unable to hash field %s", e.Field)

}
