package valueobject

type Email interface {
	IsValid() bool
	Value() string
}

type emailVO struct {
	value string
}

func NewEmail(email string) Email {
	return &emailVO{value: email}
}

func (em *emailVO) IsValid() bool {
	return false
}

func (em *emailVO) Value() string {
	return em.value
}
