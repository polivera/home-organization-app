package fakers

import (
	"github.com/go-faker/faker/v4"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
)

func HouseholdEntityFakerRandom() entity.HouseholdEntity {
	return entity.HouseholdEntity{
		Id:    uint64(faker.RandomUnixTime()),
		Name:  faker.Name(),
		Owner: uint64(faker.RandomUnixTime()),
	}
}
