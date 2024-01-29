package utils

import (
	"strings"
	"unicode"
)

func HasLower(val string, len int) bool {
	lowerFound := false
	index := 0
	for index < len && !lowerFound {
		lowerFound = unicode.IsLower(rune(val[index]))
		index++
	}
	return lowerFound
}

func HasUpper(val string, len int) bool {
	upperFound := false
	index := 0
	for index < len && !upperFound {
		upperFound = unicode.IsUpper(rune(val[index]))
		index++
	}
	return upperFound
}

func HasPasswordSymbol(val string) bool {
	return strings.ContainsAny(val, "!@$%^&*()-=_+,./<>?;':\"[]{}")
}
