package service_test

import (
	"errors"
	"testing"

	"github.com/polivera/home-organization-app/internal/household/domain"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
	"github.com/polivera/home-organization-app/internal/household/domain/service"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
	"github.com/polivera/home-organization-app/test/common/matchers"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserHouseholdsService_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockHouseholdUsersRepository(ctrl)

	t.Run("Invalid user ID", func(t *testing.T) {
		cmd := command.NewGetUserHouseholdsCommand(0)
		svc := service.NewGetUserHouseholdsService(mockRepo)
		res, err := svc.Handle(cmd)
		assert.Nil(t, res)
		assert.EqualError(t, err, "user_id is not valid")
	})

	t.Run("Repository error getting households", func(t *testing.T) {
		cmd := command.NewGetUserHouseholdsCommand(10)
		mockRepo.
			EXPECT().
			GetUserHouseholds(matchers.IDMatcher(10)).
			Times(1).
			Return(nil, errors.New("mock db error"))

		svc := service.NewGetUserHouseholdsService(mockRepo)
		res, err := svc.Handle(cmd)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Unexpected repository error: mock db error")
	})

	t.Run("Empty result getting households", func(t *testing.T) {
		cmd := command.NewGetUserHouseholdsCommand(10)
		mockRepo.
			EXPECT().
			GetUserHouseholds(matchers.IDMatcher(10)).
			Times(1).
			Return(nil, nil)

		svc := service.NewGetUserHouseholdsService(mockRepo)
		res, err := svc.Handle(cmd)
		assert.Nil(t, err)
		assert.Equal(t, []domain.HouseholdDTO(nil), res.Owned)
		assert.Equal(t, []domain.HouseholdDTO(nil), res.Participate)
	})

	t.Run("Happy path", func(t *testing.T) {
		cmd := command.NewGetUserHouseholdsCommand(10)
		mockRepo.
			EXPECT().
			GetUserHouseholds(matchers.IDMatcher(10)).
			Times(1).
			Return([]entity.Household{
				{
					ID:    1,
					Name:  "Test Owned",
					Owner: 10,
				},
				{
					ID:    2,
					Name:  "Test Participant",
					Owner: 11,
				},
			}, nil)

		svc := service.NewGetUserHouseholdsService(mockRepo)
		res, err := svc.Handle(cmd)
		assert.Nil(t, err)
		assert.Len(t, res.Owned, 1)
		assert.Len(t, res.Participate, 1)
		assert.Equal(t, "Test Owned", res.Owned[0].Name)
		assert.Equal(t, "Test Participant", res.Participate[0].Name)
	})
}
