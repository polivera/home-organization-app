package valueobject

import (
	"github.com/polivera/home-organization-app/internal/common/valueobject"
	"github.com/polivera/home-organization-app/pkg/utils"
)

type password struct {
	value string
}

type PlainPassword valueobject.ValueObject[string, PlainPassword]

func NewPlainPassword(plainPassword string) PlainPassword {
	return password{value: plainPassword}
}

func (pass password) IsValid() bool {
	strLen := len(pass.value)
	return strLen >= 8 &&
		utils.HasLower(pass.value, strLen) &&
		utils.HasUpper(pass.value, strLen) &&
		utils.HasPasswordSymbol(pass.value)
}

func (pass password) Value() string {
	return pass.value
}

func (pass password) IsEqual(vo PlainPassword) bool {
	return pass.value == vo.Value()
}

func (pass password) String() string {
	return "********"
}
