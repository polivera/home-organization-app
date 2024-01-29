package valueobject

import (
	"github.com/polivera/home-organization-app/pkg/utils"
)

var (
	passwordSpecialChars = [28]string{
		"!", "@", "$", "%", "^", "&", "*", "(", ")",
		"-", "=", "_", "+", ",", ".", "/", "<", ">",
		"?", ";", "'", ":", "\\", "\"", "[", "]", "{", "}"}
)

type PlainPassword interface {
	IsValid() bool
	GetValue() string
}

type password struct {
	value string
}

func NewPlainPassword(plainPassword string) PlainPassword {
	return &password{value: plainPassword}
}

func (pass *password) IsValid() bool {
	strLen := len(pass.value)
	return strLen >= 8 &&
		utils.HasLower(pass.value, strLen) &&
		utils.HasUpper(pass.value, strLen) &&
		utils.HasPasswordSymbol(pass.value)
}

func (pass *password) GetValue() string {
	return pass.value
}
