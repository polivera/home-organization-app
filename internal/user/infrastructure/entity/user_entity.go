package entity

const (
	StatusCreated = iota
	StatusVerified
	StatusDisabled
)

type UserEntity struct {
	Id           uint64
	Email        string
	Password     string
	Username     string
	SessionToken string
	Status       uint8
}
