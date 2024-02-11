//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package repository

import (
	commonValueObject "github.com/polivera/home-organization-app/internal/common/valueobject"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/entity"
)

type UserRepository interface {
	GetVerifiedUserByEmail(email commonValueObject.Email) (*entity.UserEntity, error)
	GetUserByEmail(email commonValueObject.Email) (*entity.UserEntity, error)
	CreateUser(userEntity *entity.UserEntity) error
}
