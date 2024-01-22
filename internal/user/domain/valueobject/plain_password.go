package valueobject

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
	return false
}

func (pass *password) GetValue() string {
	return pass.value
}
