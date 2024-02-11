package fakers

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/entity"
)

func UserEntityFakerRandom() entity.UserEntity {
	return entity.UserEntity{
		Id:           uint64(faker.RandomUnixTime()),
		Email:        faker.Email(),
		Password:     faker.Password(),
		Username:     fmt.Sprintf("Faker%s", faker.Username()),
		SessionToken: faker.JWT,
		Status:       entity.StatusCreated,
	}
}
