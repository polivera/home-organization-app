//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package repository

import (
	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/entity"
)

type UserRepository interface {
	GetVerifiedUserByEmail(email valueobject.Email) (*entity.UserEntity, error)
}
