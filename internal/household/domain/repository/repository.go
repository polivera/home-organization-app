//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package repository

import (
	commonValueObject "github.com/polivera/home-organization-app/internal/common/valueobject"
	"github.com/polivera/home-organization-app/internal/household/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
)

type HouseholdRepository interface {
	CreateHousehold(householdEntity *entity.HouseholdEntity) error
	GetUserHouseholdByName(
		name valueobject.HouseholdName,
		owner commonValueObject.ID,
	) (*entity.HouseholdEntity, error)
}