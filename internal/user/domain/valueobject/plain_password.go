package valueobject

import (
	"github.com/polivera/home-organization-app/internal/common/valueobject"
	"github.com/polivera/home-organization-app/pkg/utils"
)

type PlainPassword interface {
	valueobject.ValueObject[string]
}

type password struct {
	value string
}

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
