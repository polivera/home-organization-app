package matchers

import (
	"fmt"
	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
	"go.uber.org/mock/gomock"
	"reflect"
)

type emailMatcher struct {
	email string
}

func EmailMatcher(email string) gomock.Matcher {
	return emailMatcher{email: email}
}

func (em emailMatcher) Matches(param interface{}) bool {
	if reflect.TypeOf(param).String() != "*valueobject.emailVO" {
		return false
	}
	emailVO, ok := param.(valueobject.Email)
	if !ok {
		return false
	}
	return emailVO.Value() == em.email

}

func (em emailMatcher) String() string {
	return fmt.Sprintf("is equal to %s", em.email)
}
