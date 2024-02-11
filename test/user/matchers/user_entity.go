package matchers

import (
	"fmt"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/entity"
	"go.uber.org/mock/gomock"
	"reflect"
)

type userEntityMatcher struct {
	email    string
	password string
	username string
	status   uint8
}

func UserEntityMatcher(email string, password string, username string, status uint8) gomock.Matcher {
	return userEntityMatcher{
		email:    email,
		password: password,
		username: username,
		status:   status,
	}
}

func (uem userEntityMatcher) Matches(param interface{}) bool {
	if reflect.TypeOf(param).String() != "*entity.UserEntity" {
		return false
	}
	userEntity, ok := param.(*entity.UserEntity)
	if !ok {
		return false
	}
	return userEntity.Email == uem.email &&
		userEntity.Password != uem.password &&
		userEntity.Username == uem.username &&
		userEntity.Status == uem.status
}

func (uem userEntityMatcher) String() string {
	return fmt.Sprintf("%s - %s - %s - %d", uem.email, uem.password, uem.username, uem.status)
}
