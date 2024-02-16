package matchers

import (
	"fmt"
	commonValueObject "github.com/polivera/home-organization-app/internal/common/valueobject"
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
	if reflect.TypeOf(param).String() != "valueobject.emailVO" {
		return false
	}
	emailVO, ok := param.(commonValueObject.Email)
	if !ok {
		return false
	}
	return emailVO.Value() == em.email
}

func (em emailMatcher) String() string {
	return fmt.Sprintf("is equal to %s", em.email)
}
