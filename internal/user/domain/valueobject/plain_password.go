package valueobject

import (
	"strings"
	"unicode"
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
		hasLower(pass.value, strLen) &&
		hasUpper(pass.value, strLen) &&
		hasSymbol(pass.value)
}

func (pass *password) GetValue() string {
	return pass.value
}

func hasLower(val string, strLen int) bool {
	lowerFound := false
	index := 0
	for index < strLen && !lowerFound {
		lowerFound = unicode.IsLower(rune(val[index]))
		index++
	}
	return lowerFound
}

func hasUpper(val string, strLen int) bool {
	upperFound := false
	index := 0
	for index < strLen && !upperFound {
		upperFound = unicode.IsUpper(rune(val[index]))
		index++
	}
	return upperFound
}

func hasSymbol(val string) bool {
	return strings.ContainsAny(val, "!@$%^&*()-=_+,./<>?;':\"[]{}")
}
