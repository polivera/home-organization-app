package service_test

import (
	"errors"
	"github.com/polivera/home-organization-app/internal/household/domain/service"
	"testing"

	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	householdRepoPkg "github.com/polivera/home-organization-app/internal/household/domain/repository"
	"github.com/polivera/home-organization-app/test/common/matchers"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAddHouseholdUserService_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHouseholdUsersRepo := householdRepoPkg.NewMockHouseholdUsersRepository(ctrl)

	t.Run("invalid household ID", func(t *testing.T) {
		cmd := command.NewAddUserToHouseholdCommand(0, 22)
		serv := service.NewAddHouseholdUserService(mockHouseholdUsersRepo)
		resp, err := serv.Handle(cmd)
		assert.EqualError(t, err, "household_id is not valid")
		assert.Nil(t, resp)
	})

	t.Run("invalid user ID", func(t *testing.T) {
		cmd := command.NewAddUserToHouseholdCommand(11, 0)
		serv := service.NewAddHouseholdUserService(mockHouseholdUsersRepo)
		resp, err := serv.Handle(cmd)
		assert.EqualError(t, err, "user_id is not valid")
		assert.Nil(t, resp)
	})

	t.Run("Error foreign key not found", func(t *testing.T) {
		cmd := command.NewAddUserToHouseholdCommand(11, 22)
		mockHouseholdUsersRepo.
			EXPECT().
			AddHouseholdUser(matchers.IDMatcher(11), matchers.IDMatcher(22)).
			Return(database.ErrorForeignRowNotFound{Table: "mock-table"})
		serv := service.NewAddHouseholdUserService(mockHouseholdUsersRepo)
		resp, err := serv.Handle(cmd)
		assert.EqualError(t, err, "Repository error: foreign row on table `mock-table` not found")
		assert.Nil(t, resp)
	})

	t.Run("Error foreign key not found", func(t *testing.T) {
		cmd := command.NewAddUserToHouseholdCommand(11, 22)
		mockHouseholdUsersRepo.
			EXPECT().
			AddHouseholdUser(matchers.IDMatcher(11), matchers.IDMatcher(22)).
			Return(database.ErrorDuplicateKey{Key: "mock-key"})
		serv := service.NewAddHouseholdUserService(mockHouseholdUsersRepo)
		resp, err := serv.Handle(cmd)
		assert.EqualError(t, err, "Repository error: record already exist under key `mock-key`")
		assert.Nil(t, resp)
	})

	t.Run("Error foreign key not found", func(t *testing.T) {
		cmd := command.NewAddUserToHouseholdCommand(11, 22)
		mockHouseholdUsersRepo.
			EXPECT().
			AddHouseholdUser(matchers.IDMatcher(11), matchers.IDMatcher(22)).
			Return(errors.New("mock error"))
		serv := service.NewAddHouseholdUserService(mockHouseholdUsersRepo)
		resp, err := serv.Handle(cmd)
		assert.EqualError(t, err, "Unexpected repository error: mock error")
		assert.Nil(t, resp)
	})

	t.Run("happy path", func(t *testing.T) {
		cmd := command.NewAddUserToHouseholdCommand(11, 22)
		mockHouseholdUsersRepo.
			EXPECT().
			AddHouseholdUser(matchers.IDMatcher(11), matchers.IDMatcher(22)).
			Return(nil)
		serv := service.NewAddHouseholdUserService(mockHouseholdUsersRepo)
		resp, err := serv.Handle(cmd)
		assert.NoError(t, err)
		assert.Equal(t, uint64(11), resp.ID)
		assert.Equal(t, uint64(22), resp.Participants[0].ID)
	})
}
