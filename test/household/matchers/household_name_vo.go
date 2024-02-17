package matchers

import (
	"fmt"
	"github.com/polivera/home-organization-app/internal/household/domain/valueobject"
	"go.uber.org/mock/gomock"
	"reflect"
)

type houseHoldMatcher struct {
	name string
}

func HouseholdMatcher(name string) gomock.Matcher {
	return houseHoldMatcher{name: name}
}

func (em houseHoldMatcher) Matches(param interface{}) bool {
	if reflect.TypeOf(param).String() != "valueobject.householdName" {
		return false
	}
	householdNameVO, ok := param.(valueobject.HouseholdNameVO)
	if !ok {
		return false
	}
	return householdNameVO.Value() == em.name
}

func (em houseHoldMatcher) String() string {
	return fmt.Sprintf("is equal to %s", em.name)
}
