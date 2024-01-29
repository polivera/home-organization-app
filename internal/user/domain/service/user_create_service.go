package service

import (
	"github.com/polivera/home-organization-app/internal/user/domain"
	"github.com/polivera/home-organization-app/internal/user/domain/command"
	"github.com/polivera/home-organization-app/internal/user/domain/repository"
)

type CreateUserService interface {
	Handle(command command.UserCreateCommand) (*domain.UserDTO, error)
}

type createUserService struct {
	userRepo repository.UserRepository
}

func NewCreateUserService(repo repository.UserRepository) CreateUserService {
	return &createUserService{userRepo: repo}
}

func (cus *createUserService) Handle(command command.UserCreateCommand) (*domain.UserDTO, error) {
	//email := valueobject.NewEmail(command.Email())
	//password := valueobject.NewPlainPassword(command.Password())

	return nil, nil
}
