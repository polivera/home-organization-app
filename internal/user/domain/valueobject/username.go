package valueobject

import (
	"github.com/polivera/home-organization-app/internal/common/valueobject"
	"regexp"
)

type Username interface {
	valueobject.ValueObject[string]
}

type usernameVO struct {
	value string
}

func NewUsername(username string) Username {
	return usernameVO{value: username}
}

func (un usernameVO) IsValid() bool {
	valLen := len(un.value)
	return valLen >= 8 &&
		regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(un.value)
}

func (un usernameVO) Value() string {
	return un.value
}
