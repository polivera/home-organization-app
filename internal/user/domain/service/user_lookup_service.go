package service

import (
	"fmt"
	"github.com/polivera/home-organization-app/internal/user/domain"
	"github.com/polivera/home-organization-app/internal/user/domain/repository"
	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
)

type LookupService interface {
	Handle(command domain.UserLookupCommand) (domain.UserDTO, error)
}

type lookupService struct {
	userRepo repository.UserRepository
}

func NewLookupService(repo repository.UserRepository) LookupService {
	return &lookupService{userRepo: repo}
}

func (ls *lookupService) Handle(command domain.UserLookupCommand) (domain.UserDTO, error) {
	entity, err := ls.userRepo.GetVerifiedUserByEmail(valueobject.NewEmail(command.Email()))
	if err != nil {
		fmt.Printf("%s", err.Error())
		return domain.UserDTO{}, err
	}
	fmt.Println(entity)
	return domain.UserDTO{}, nil
}
