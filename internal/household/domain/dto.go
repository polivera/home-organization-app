package domain

type HouseholdDTO struct {
	ID           uint64
	Name         string
	OwnerID      uint64
	Owner        Participant
	Participants []Participant
}

type Participant struct {
	ID   uint64
	Name string
}

type UserHouseholdsDTO struct {
	UserID      uint64
	Owned       []HouseholdDTO
	Participate []HouseholdDTO
}
