package valueobject

type IDVO struct {
	value uint64
}

type ID interface {
	ValueObject[uint64]
}

func NewID(id uint64) ID {
	return IDVO{value: id}
}

func (id IDVO) Value() uint64 {
	return id.value
}

func (id IDVO) IsValid() bool {
	return id.value > 0
}
