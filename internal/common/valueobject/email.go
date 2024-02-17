package valueobject

import (
	"fmt"
	"regexp"
	"strings"
)

// @see https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address
var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9.!#$%&'*+/=?^_\x60{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`,
)

type EmailVO ValueObject[string, EmailVO]

type emailVO struct {
	value string
}

func NewEmail(email string) EmailVO {
	return emailVO{value: email}
}

func (em emailVO) IsValid() bool {
	return emailRegex.MatchString(strings.ToLower(em.value))
}

func (em emailVO) Value() string {
	return em.value
}

func (em emailVO) IsEqual(vo EmailVO) bool {
	return em.value == vo.Value()
}

func (em emailVO) String() string {
	return fmt.Sprintf("email address: %s", em.value)
}
