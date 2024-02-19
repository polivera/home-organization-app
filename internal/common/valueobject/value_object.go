package valueobject

type ValueObject[TP any, IN any] interface {
	Value() TP
	IsValid() bool
	IsEqual(vo IN) bool
	String() string
}
