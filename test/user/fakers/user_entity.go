package fakers

import (
	"github.com/go-faker/faker/v4"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/entity"
)

func UserFakerEntityRandom() entity.UserEntity {
	return entity.UserEntity{
		Id:           uint64(faker.RandomUnixTime()),
		Email:        faker.Email(),
		Password:     faker.Password(),
		Username:     faker.Username(),
		SessionToken: faker.JWT,
		Status:       entity.StatusCreated,
	}
}
