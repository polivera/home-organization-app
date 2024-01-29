package service

import (
	"github.com/polivera/home-organization-app/internal/common"
	"github.com/polivera/home-organization-app/internal/user/domain"
	"github.com/polivera/home-organization-app/internal/user/domain/command"
	"github.com/polivera/home-organization-app/internal/user/domain/repository"
	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
)

type LookupService interface {
	Handle(command command.UserLookupCommand) (*domain.UserDTO, error)
}

type lookupService struct {
	userRepo repository.UserRepository
}

func NewLookupService(repo repository.UserRepository) LookupService {
	return &lookupService{userRepo: repo}
}

func (ls *lookupService) Handle(command command.UserLookupCommand) (*domain.UserDTO, error) {
	email := valueobject.NewEmail(command.Email())
	password := valueobject.NewPlainPassword(command.Password())

	if !email.IsValid() {
		return nil, common.ErrorValidation{Field: "email"}
	}

	if !password.IsValid() {
		return nil, common.ErrorValidation{Field: "password"}
	}

	entity, err := ls.userRepo.GetVerifiedUserByEmail(email)
	if err != nil {
		return nil, common.ErrorNotFound{Item: email.Value()}
	}

	hashPass := valueobject.NewHashPassword(entity.Password)
	if !hashPass.IsPasswordValid(password) {
		return nil, common.ErrorNotFound{Item: email.Value()}
	}

	return &domain.UserDTO{
		Id:       entity.Id,
		Email:    entity.Email,
		Username: entity.Username,
	}, nil
}
