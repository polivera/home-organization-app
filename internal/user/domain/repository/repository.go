//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package repository

import (
	commonValueObject "github.com/polivera/home-organization-app/internal/common/valueobject"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/entity"
)

type UserRepository interface {
	GetVerifiedUserByEmail(email commonValueObject.EmailVO) (*entity.UserEntity, error)
	GetUserByEmail(email commonValueObject.EmailVO) (*entity.UserEntity, error)
	CreateUser(userEntity *entity.UserEntity) error
	GetUserByID(id commonValueObject.IDVO) (*entity.UserEntity, error)
}
