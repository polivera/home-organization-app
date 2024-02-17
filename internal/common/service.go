package common

type DomainService[CM any, DTO any] interface {
	Handle(command CM) (*DTO, error)
}
