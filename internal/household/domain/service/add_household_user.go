package service

import (
	"errors"
	"github.com/polivera/home-organization-app/internal/common"
	"github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/household/domain"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
	userRepository "github.com/polivera/home-organization-app/internal/user/domain/repository"
)

type addHouseholdUserService struct {
	householdRepo     repository.HouseholdRepository
	householdUserRepo repository.HouseholdUsersRepository
	userRepo          userRepository.UserRepository
}

func NewAddHouseholdUserService(
	householdRepo repository.HouseholdRepository,
	householdUserRepo repository.HouseholdUsersRepository,
	userRepo userRepository.UserRepository,
) common.DomainService[
	command.AddUserToHouseholdCommand,
	domain.HouseholdDTO,
] {
	return &addHouseholdUserService{
		householdRepo:     householdRepo,
		householdUserRepo: householdUserRepo,
		userRepo:          userRepo,
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
		Id:           householdID.Value(),
		Name:         "",
		Owner:        0,
		Participants: []domain.Participant{{ID: userID.Value()}},
	}, nil
}
