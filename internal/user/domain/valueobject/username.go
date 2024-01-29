package valueobject

type Username interface {
	IsValid() bool
	Value() string
}

type usernameVO struct {
	value string
}

func NewUsername(username string) Username {
	return &usernameVO{value: username}
}

func (un *usernameVO) IsValid() bool {
	valLen := len(un.value)
	// todo: check that only has alpha-numeric chars
	return valLen >= 8
}

func (un *usernameVO) Value() string {
	return un.value
}
