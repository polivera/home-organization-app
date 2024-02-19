package service

import (
	"github.com/polivera/home-organization-app/internal/common"
	commonValueObject "github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/user"
	"github.com/polivera/home-organization-app/internal/user/domain"
	"github.com/polivera/home-organization-app/internal/user/domain/command"
	"github.com/polivera/home-organization-app/internal/user/domain/repository"
	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/entity"
)

type createUserService struct {
	userRepo repository.UserRepository
}

func NewCreateUserService(repo repository.UserRepository) common.DomainService[
	command.UserCreateCommand,
	domain.UserDTO,
] {
	return &createUserService{userRepo: repo}
}

func (cus *createUserService) Handle(command command.UserCreateCommand) (*domain.UserDTO, error) {
	email := commonValueObject.NewEmail(command.Email())
	password := valueobject.NewPlainPassword(command.Password())
	username := valueobject.NewUsername(command.Username())

	if !email.IsValid() {
		return nil, common.ErrorValidation{Field: "email"}
	}
	if !password.IsValid() {
		return nil, common.ErrorValidation{Field: "password"}
	}
	if !username.IsValid() {
		return nil, common.ErrorValidation{Field: "username"}
	}

	existingUser, err := cus.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, common.RepositoryUnexpectedError{Message: err.Error()}
	}
	if existingUser != nil {
		return nil, user.ErrorUserExist{Email: email.Value()}
	}

	hashPassword, err := valueobject.NewHashFromPlain(password)
	if err != nil {
		return nil, user.ErrorCantHashField{Field: "password"}
	}

	var userEntity entity.UserEntity
	userEntity.Email = email.Value()
	userEntity.Password = hashPassword.Value()
	userEntity.Username = username.Value()
	userEntity.Status = entity.StatusCreated
	err = cus.userRepo.CreateUser(&userEntity)
	if err != nil {
		return nil, common.RepositoryUnexpectedError{Message: err.Error()}
	}

	return &domain.UserDTO{
		Id:       userEntity.Id,
		Email:    userEntity.Email,
		Username: userEntity.Username,
	}, nil
}
