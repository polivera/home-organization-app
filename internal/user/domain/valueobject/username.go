package valueobject

import (
	"fmt"
	"regexp"

	"github.com/polivera/home-organization-app/internal/common/domain/valueobject"
)

type UsernameVO valueobject.ValueObject[string, UsernameVO]

type usernameVO struct {
	value string
}

func NewUsername(username string) UsernameVO {
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

func (un usernameVO) IsEqual(vo UsernameVO) bool {
	return un.value == vo.Value()
}

func (un usernameVO) String() string {
	return fmt.Sprintf("username: %s", un.value)
}
