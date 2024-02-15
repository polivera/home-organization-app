package valueobject

import (
	"regexp"
	"strings"
)

// @see https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address
var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9.!#$%&'*+/=?^_\x60{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`,
)

type Email interface {
	IsValid() bool
	Value() string
}

type emailVO struct {
	value string
}

func NewEmail(email string) Email {
	return emailVO{value: email}
}

func (em emailVO) IsValid() bool {
	return emailRegex.MatchString(strings.ToLower(em.value))
}

func (em emailVO) Value() string {
	return em.value
}
