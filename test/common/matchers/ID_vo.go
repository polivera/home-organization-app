package matchers

import (
	"fmt"
	commonValueObject "github.com/polivera/home-organization-app/internal/common/valueobject"
	"go.uber.org/mock/gomock"
	"reflect"
)

type idMatcher struct {
	id uint64
}

func IDMatcher(id uint64) gomock.Matcher {
	return idMatcher{id: id}
}

func (em idMatcher) Matches(param interface{}) bool {
	if reflect.TypeOf(param).String() != "valueobject.idVO" {
		return false
	}
	idVO, ok := param.(commonValueObject.IDVO)
	if !ok {
		return false
	}
	return idVO.Value() == em.id
}

func (em idMatcher) String() string {
	return fmt.Sprintf("is equal to %d", em.id)
}
