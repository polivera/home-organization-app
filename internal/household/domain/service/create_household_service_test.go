package service

import (
	"errors"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
	commonMatchers "github.com/polivera/home-organization-app/test/common/matchers"
	"github.com/polivera/home-organization-app/test/household/fakers"
	"github.com/polivera/home-organization-app/test/household/matchers"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCreateHouseholdService_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	householdRepo := repository.NewMockHouseholdRepository(ctrl)

	t.Run("fail if household is invalid", func(t *testing.T) {
		cmd := command.NewCreateHouseholdCommand("", 10)
		handle := NewCreateHouseholdService(householdRepo)
		_, err := handle.Handle(cmd)
		assert.EqualError(t, err, "household name is not valid")
	})

	t.Run("fail if owner is invalid", func(t *testing.T) {
		cmd := command.NewCreateHouseholdCommand("household", 0)
		handle := NewCreateHouseholdService(householdRepo)
		_, err := handle.Handle(cmd)
		assert.EqualError(t, err, "owner is not valid")
	})

	t.Run("fail db error on select", func(t *testing.T) {
		cmd := command.NewCreateHouseholdCommand("household", 1)
		householdRepo.EXPECT().
			GetUserHouseholdByName(matchers.HouseholdMatcher("household"), commonMatchers.IDMatcher(1)).
			Return(nil, errors.New("some-mock-error"))
		handle := NewCreateHouseholdService(householdRepo)
		_, err := handle.Handle(cmd)
		assert.EqualError(t, err, "Unexpected repository error: some-mock-error")
	})

	t.Run("fail if household exist for user", func(t *testing.T) {
		cmd := command.NewCreateHouseholdCommand("THE HOLD", 25)
		entity := fakers.HouseholdEntityFakerRandom()
		householdRepo.EXPECT().
			GetUserHouseholdByName(matchers.HouseholdMatcher("THE HOLD"), commonMatchers.IDMatcher(25)).
			Return(&entity, nil)
		handle := NewCreateHouseholdService(householdRepo)
		_, err := handle.Handle(cmd)
		assert.EqualError(t, err, "Household named THE HOLD already exist for user 25")
	})

	t.Run("fail if unexpected error creating household", func(t *testing.T) {
		cmd := command.NewCreateHouseholdCommand("THE HOLD", 25)
		householdRepo.EXPECT().
			GetUserHouseholdByName(matchers.HouseholdMatcher("THE HOLD"), commonMatchers.IDMatcher(25)).
			Times(1).
			Return(nil, nil)
		householdRepo.EXPECT().
			CreateHousehold(matchers.NewHouseholdEntityMatcher(0, "THE HOLD", 25)).
			Times(1).
			Return(errors.New("another-mock-error"))

		handle := NewCreateHouseholdService(householdRepo)
		_, err := handle.Handle(cmd)
		assert.EqualError(t, err, "Unexpected repository error: another-mock-error")
	})

	t.Run("happy path", func(t *testing.T) {
		cmd := command.NewCreateHouseholdCommand("THE HOLD", 25)
		householdRepo.EXPECT().
			GetUserHouseholdByName(matchers.HouseholdMatcher("THE HOLD"), commonMatchers.IDMatcher(25)).
			Times(1).
			Return(nil, nil)
		householdRepo.EXPECT().
			CreateHousehold(matchers.NewHouseholdEntityMatcher(0, "THE HOLD", 25)).
			Times(1).
			DoAndReturn(func(hhEntity *entity.HouseholdEntity) error {
				hhEntity.Id = 123
				return nil
			})

		handle := NewCreateHouseholdService(householdRepo)
		householdDTO, err := handle.Handle(cmd)
		expectedEntity := entity.HouseholdEntity{
			Id:    123,
			Name:  "THE HOLD",
			Owner: 25,
		}
		assert.NoError(t, err)
		assert.Equal(t, householdDTO.Id, expectedEntity.Id)
		assert.Equal(t, householdDTO.Name, expectedEntity.Name)
		assert.Equal(t, householdDTO.Owner, expectedEntity.Owner)
	})
}
