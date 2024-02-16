package valueobject

import "github.com/polivera/home-organization-app/internal/common/valueobject"

type HouseholdName interface {
	valueobject.ValueObject[string]
}

type householdName struct {
	value string
}

func NewHouseholdName(name string) HouseholdName {
	return householdName{value: name}
}

func (hn householdName) Value() string {
	return hn.value
}
func (hn householdName) IsValid() bool {
	return len(hn.value) > 0
}
