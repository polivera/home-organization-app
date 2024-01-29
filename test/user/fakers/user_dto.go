package fakers

import (
	"github.com/go-faker/faker/v4"
	"github.com/polivera/home-organization-app/internal/user/domain"
)

func UserDTOFakerRandom() domain.UserDTO {
	return domain.UserDTO{
		Email:    faker.Email(),
		Username: faker.Username(),
	}
}
