package valueobject

type ValueObject[TP any] interface {
	Value() TP
	IsValid() bool
}
