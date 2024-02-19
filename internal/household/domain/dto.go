package domain

type HouseholdDTO struct {
	ID           uint64
	Name         string
	Owner        uint64
	Participants []Participant
}

type Participant struct {
	ID   uint64
	Name string
}
