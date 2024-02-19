package service

import (
	"errors"

	"github.com/polivera/home-organization-app/internal/common"
	"github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/household/domain"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
)

type addHouseholdUserService struct {
	householdUserRepo repository.HouseholdUsersRepository
}

func NewAddHouseholdUserService(
	householdUserRepo repository.HouseholdUsersRepository,
) common.DomainService[
	command.AddUserToHouseholdCommand,
	domain.HouseholdDTO,
] {
	return &addHouseholdUserService{
		householdUserRepo: householdUserRepo,
	}
}

func (ahs addHouseholdUserService) Handle(command command.AddUserToHouseholdCommand) (*domain.HouseholdDTO, error) {
	householdID := valueobject.NewID(command.Household())
	userID := valueobject.NewID(command.User())

	if !householdID.IsValid() {
		return nil, common.ErrorValidation{Field: "household_id"}
	}
	if !userID.IsValid() {
		return nil, common.ErrorValidation{Field: "user_id"}
	}

	err := ahs.householdUserRepo.AddHouseholdUser(householdID, userID)
	if err != nil {
		if errors.As(err, database.ErrForeignRowNotFound) || errors.As(err, database.ErrDuplicateKey) {
			return nil, common.RepositoryError{Message: err.Error()}
		}
		return nil, common.RepositoryUnexpectedError{Message: err.Error()}
	}

	return &domain.HouseholdDTO{
		ID:           householdID.Value(),
		Name:         "",
		Owner:        0,
		Participants: []domain.Participant{{ID: userID.Value()}},
	}, nil
}
