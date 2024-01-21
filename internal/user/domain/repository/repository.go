package repository

import "github.com/polivera/home-organization-app/internal/user/infrastructure/entity"

type UserRepository interface {
	GetVerifiedUserByEmail(email string) (entity.UserEntity, error)
}
