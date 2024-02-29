package entity

type Household struct {
	ID    uint64
	Name  string
	Owner uint64
}

type HouseholdUser struct {
	HouseholdID uint64
	UserID      uint64
}
