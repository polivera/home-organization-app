//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package repository

import (
	commonValueObject "github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/household/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
)

type HouseholdRepository interface {
	CreateHousehold(householdEntity *entity.Household) error
	GetUserHouseholdByName(
		name valueobject.HouseholdNameVO,
		owner commonValueObject.IDVO,
	) (*entity.Household, error)
	GetHouseholdByID(id commonValueObject.IDVO) (*entity.Household, error)
	GetHouseholdsByUserID(id commonValueObject.IDVO) ()
}

type HouseholdUsersRepository interface {
	AddHouseholdUser(householdID commonValueObject.IDVO, userID commonValueObject.IDVO) error
}
