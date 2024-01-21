package repository

type UserRepository interface {
	GetActiveUserByEmail(email string)
}
