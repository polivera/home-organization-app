package service

import (
	"github.com/polivera/home-organization-app/internal/common"
	"github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/household/domain"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
)

type getUserHouseholdsService struct {
	householdUserRepo repository.HouseholdUsersRepository
}

func NewGetUserHouseholdsService(
	householdUserRepo repository.HouseholdUsersRepository,
) common.DomainService[
	command.GetUserHouseholdsCommand,
	domain.HouseholdDTO,
] {
	return &getUserHouseholdsService{
		householdUserRepo: householdUserRepo,
	}
}

func (guh getUserHouseholdsService) Handle(command command.GetUserHouseholdsCommand) (*domain.HouseholdDTO, error) {
	userID := valueobject.NewID(command.User())

	if !userID.IsValid() {
		return nil, common.ErrorValidation{Field: "user_id"}
	}



	return nil, nil
}
