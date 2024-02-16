//go:build unit

package service_test

import (
	"errors"
	"testing"

	"github.com/polivera/home-organization-app/internal/user/domain/command"
	"github.com/polivera/home-organization-app/internal/user/domain/repository"
	"github.com/polivera/home-organization-app/internal/user/domain/service"
	commonMatchers "github.com/polivera/home-organization-app/test/common/matchers"
	"github.com/polivera/home-organization-app/test/user/fakers"
	"github.com/polivera/home-organization-app/test/user/matchers"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepoMock := repository.NewMockUserRepository(ctrl)

	t.Run("fail if invalid email", func(t *testing.T) {
		cmd := command.NewUserCreateCommand("invalid-email", "Test.123", "TestUser")
		handler := service.NewCreateUserService(userRepoMock)
		_, err := handler.Handle(cmd)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email")
	})

	t.Run("fail if password is invalid", func(t *testing.T) {
		cmd := command.NewUserCreateCommand("test@test.local", "wrongpass", "TestUser")
		handler := service.NewCreateUserService(userRepoMock)
		_, err := handler.Handle(cmd)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password")
	})

	t.Run("fail if username is invalid", func(t *testing.T) {
		cmd := command.NewUserCreateCommand("test@test.local", "Test.123", "us@!#@#r")
		handler := service.NewCreateUserService(userRepoMock)
		_, err := handler.Handle(cmd)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "username")
	})

	t.Run("unexpected database error", func(t *testing.T) {
		cmd := command.NewUserCreateCommand("test@test.local", "Test.123", "TestUser")
		userRepoMock.EXPECT().
			GetUserByEmail(commonMatchers.EmailMatcher("test@test.local")).
			Times(1).
			Return(nil, errors.New("database error"))
		handler := service.NewCreateUserService(userRepoMock)
		_, err := handler.Handle(cmd)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
	})

	t.Run("fail if the user exist", func(t *testing.T) {
		cmd := command.NewUserCreateCommand("test@test.local", "Test.123", "TestUser")
		mockUserEntity := fakers.UserEntityFakerRandom()
		userRepoMock.EXPECT().
			GetUserByEmail(commonMatchers.EmailMatcher("test@test.local")).
			Times(1).
			Return(&mockUserEntity, nil)
		handler := service.NewCreateUserService(userRepoMock)
		_, err := handler.Handle(cmd)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email")
	})

	t.Run("error storing user", func(t *testing.T) {
		mockUserEntity := fakers.UserEntityFakerRandom()
		cmd := command.NewUserCreateCommand(mockUserEntity.Email, "Test.123", mockUserEntity.Username)
		userRepoMock.EXPECT().
			GetUserByEmail(commonMatchers.EmailMatcher(mockUserEntity.Email)).
			Return(nil, nil)
		userRepoMock.EXPECT().
			CreateUser(
				matchers.UserEntityMatcher(
					mockUserEntity.Email,
					mockUserEntity.Password,
					mockUserEntity.Username,
					mockUserEntity.Status,
				),
			).
			Times(1).
			Return(errors.New("error storing user"))
		handler := service.NewCreateUserService(userRepoMock)
		_, err := handler.Handle(cmd)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error storing user")
	})

}
