package valueobject

import "fmt"

type IDVO ValueObject[uint64, IDVO]

type idVO struct {
	value uint64
}

func NewID(id uint64) IDVO {
	return idVO{value: id}
}

func (id idVO) Value() uint64 {
	return id.value
}

func (id idVO) IsValid() bool {
	return id.value > 0
}

func (id idVO) IsEqual(vo IDVO) bool {
	return id.value == vo.Value()
}

func (id idVO) String() string {
	return fmt.Sprintf("ID: %d", id.value)
}
