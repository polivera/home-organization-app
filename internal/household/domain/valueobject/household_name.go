package valueobject

import (
	"fmt"
	"github.com/polivera/home-organization-app/internal/common/domain/valueobject"
)

type HouseholdNameVO valueobject.ValueObject[string, HouseholdNameVO]

type householdName struct {
	value string
}

func NewHouseholdName(name string) valueobject.ValueObject[string, HouseholdNameVO] {
	return householdName{value: name}
}

func (hn householdName) Value() string {
	return hn.value
}

func (hn householdName) IsValid() bool {
	return len(hn.value) > 0
}

func (hn householdName) IsEqual(vo HouseholdNameVO) bool {
	return hn.value == vo.Value()
}

func (hn householdName) String() string {
	return fmt.Sprintf("household name: %s", hn.value)
}
