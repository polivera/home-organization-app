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
	domain.UserHouseholdsDTO,
] {
	return &getUserHouseholdsService{
		householdUserRepo: householdUserRepo,
	}
}

func (guh getUserHouseholdsService) Handle(command command.GetUserHouseholdsCommand) (*domain.UserHouseholdsDTO, error) {
	userID := valueobject.NewID(command.User())

	if !userID.IsValid() {
		return nil, common.ErrorValidation{Field: "user_id"}
	}

	res, err := guh.householdUserRepo.GetUserHouseholds(userID)

	if err != nil {
		return nil, common.RepositoryUnexpectedError{Message: err.Error()}
	}

	var ownedHouseholds []domain.HouseholdDTO
	var participantHousehold []domain.HouseholdDTO
	var householdDTO domain.HouseholdDTO
	for _, household := range res {
		householdDTO.ID = household.ID
		householdDTO.OwnerID = household.Owner
		householdDTO.Name = household.Name

		if household.Owner == userID.Value() {
			ownedHouseholds = append(ownedHouseholds, householdDTO)
			continue
		}

		participantHousehold = append(participantHousehold, householdDTO)
	}

	return &domain.UserHouseholdsDTO{
		UserID:      userID.Value(),
		Owned:       ownedHouseholds,
		Participate: participantHousehold,
	}, nil
}
