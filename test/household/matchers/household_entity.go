package matchers

import (
	"fmt"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
	"go.uber.org/mock/gomock"
	"reflect"
)

type householdEntityMatcher struct {
	id    uint64
	name  string
	owner uint64
}

func NewHouseholdEntityMatcher(id uint64, name string, owner uint64) gomock.Matcher {
	return householdEntityMatcher{
		id:    id,
		name:  name,
		owner: owner,
	}
}

func (hem householdEntityMatcher) Matches(param interface{}) bool {
	if reflect.TypeOf(param).String() != "*entity.Household" {
		return false
	}
	householdEntity, ok := param.(*entity.Household)
	if !ok {
		return false
	}
	return householdEntity.Id == hem.id &&
		householdEntity.Name == hem.name &&
		householdEntity.Owner == hem.owner
}

func (hem householdEntityMatcher) String() string {
	return fmt.Sprintf("%d - %s - %d", hem.id, hem.name, hem.owner)
}
