package service_test

import (
	"errors"
	"testing"

	"github.com/polivera/home-organization-app/internal/common"
	"github.com/polivera/home-organization-app/internal/user/domain/command"
	"github.com/polivera/home-organization-app/internal/user/domain/repository"
	"github.com/polivera/home-organization-app/internal/user/domain/service"
	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
	commonMatchers "github.com/polivera/home-organization-app/test/common"
	"github.com/polivera/home-organization-app/test/user/fakers"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLookupService_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepository := repository.NewMockUserRepository(ctrl)

	t.Run("Email not valid", func(t *testing.T) {
		cmd := command.NewUserLookupCommand("wrongemail", "Test.123")
		lookupService := service.NewLookupService(mockUserRepository)

		dto, err := lookupService.Handle(cmd)
		assert.Nil(t, dto)
		assert.Error(t, err)

		assert.IsType(t, common.ErrorValidation{}, err)
		assert.Equal(t, "email is not valid", err.Error())
	})

	t.Run("Password not valid", func(t *testing.T) {

		cmd := command.NewUserLookupCommand("valid@emial.local", "wrongpass")
		lookupService := service.NewLookupService(mockUserRepository)

		dto, err := lookupService.Handle(cmd)
		assert.Nil(t, dto)
		assert.Error(t, err)

		assert.IsType(t, common.ErrorValidation{}, err)
		assert.Equal(t, "password is not valid", err.Error())
	})

	t.Run("Fail to retrieve user", func(t *testing.T) {
		validEmail := "valid@email.local"
		cmd := command.NewUserLookupCommand(validEmail, "Test.123")

		mockUserRepository.
			EXPECT().
			GetVerifiedUserByEmail(commonMatchers.EmailMatcher(validEmail)).
			Times(1).
			Return(nil, errors.New("db error"))

		lookupService := service.NewLookupService(mockUserRepository)
		dto, err := lookupService.Handle(cmd)
		assert.Nil(t, dto)
		assert.Error(t, err)

		assert.IsType(t, common.ErrorNotFound{}, err)
		assert.Equal(t, validEmail+" not found", err.Error())
	})

	t.Run("Fail to retrieve user", func(t *testing.T) {
		validEmail := "valid@email.local"
		cmd := command.NewUserLookupCommand(validEmail, "Test.123")
		fakeEntity := fakers.UserEntityFakerRandom()

		mockUserRepository.
			EXPECT().
			GetVerifiedUserByEmail(commonMatchers.EmailMatcher(validEmail)).
			Times(1).
			Return(&fakeEntity, nil)

		lookupService := service.NewLookupService(mockUserRepository)
		dto, err := lookupService.Handle(cmd)
		assert.Nil(t, dto)
		assert.Error(t, err)

		assert.IsType(t, common.ErrorNotFound{}, err)
		assert.Equal(t, validEmail+" not found", err.Error())
	})

	t.Run("Happy Path", func(t *testing.T) {
		validEmail := "valid@email.local"
		validPass := "Test.123"
		cmd := command.NewUserLookupCommand(validEmail, validPass)
		fakeEntity := fakers.UserEntityFakerRandom()
		hashVO, err := valueobject.NewHashFromPlain(valueobject.NewPlainPassword(validPass))
		assert.NoError(t, err)

		fakeEntity.Password = hashVO.Value()
		mockUserRepository.
			EXPECT().
			GetVerifiedUserByEmail(commonMatchers.EmailMatcher(validEmail)).
			Times(1).
			Return(&fakeEntity, nil)

		lookupService := service.NewLookupService(mockUserRepository)
		dto, err := lookupService.Handle(cmd)
		assert.NoError(t, err)
		assert.Equal(t, fakeEntity.Email, dto.Email)
		assert.Equal(t, fakeEntity.Username, dto.Username)
	})
}
