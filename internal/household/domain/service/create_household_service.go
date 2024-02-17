package service

import (
	"github.com/polivera/home-organization-app/internal/common"
	commonVO "github.com/polivera/home-organization-app/internal/common/valueobject"
	"github.com/polivera/home-organization-app/internal/household"
	"github.com/polivera/home-organization-app/internal/household/domain"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
	"github.com/polivera/home-organization-app/internal/household/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
)

type createHouseholdService struct {
	householdRepo repository.HouseholdRepository
}

func NewCreateHouseholdService(repo repository.HouseholdRepository) common.DomainService[
	command.CreateHouseholdCommand,
	domain.HouseholdDTO,
] {
	return &createHouseholdService{
		householdRepo: repo,
	}
}

func (chs *createHouseholdService) Handle(command command.CreateHouseholdCommand) (*domain.HouseholdDTO, error) {
	hhName := valueobject.NewHouseholdName(command.Name())
	owner := commonVO.NewID(command.Owner())

	if !hhName.IsValid() {
		return nil, common.ErrorValidation{Field: "household name"}
	}
	if !owner.IsValid() {
		return nil, common.ErrorValidation{Field: "owner"}
	}

	existingHousehold, err := chs.householdRepo.GetUserHouseholdByName(hhName, owner)
	if err != nil {
		return nil, common.RepositoryUnexpectedError{Message: err.Error()}
	}
	if existingHousehold != nil {
		return nil, household.ErrorHouseholdExist{
			Name:  hhName.Value(),
			Owner: owner.Value(),
		}
	}

	var hhEntity entity.HouseholdEntity
	hhEntity.Name = hhName.Value()
	hhEntity.Owner = owner.Value()
	if err = chs.householdRepo.CreateHousehold(&hhEntity); err != nil {
		return nil, common.RepositoryUnexpectedError{Message: err.Error()}
	}

	return &domain.HouseholdDTO{
		Id:    hhEntity.Id,
		Name:  hhEntity.Name,
		Owner: hhEntity.Owner,
	}, nil
}
