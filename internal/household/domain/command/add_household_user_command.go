package command

type AddUserToHouseholdCommand struct {
	householdID uint64
	userID      uint64
}

func NewAddUserToHouseholdCommand(householdID uint64, userID uint64) AddUserToHouseholdCommand {
	return AddUserToHouseholdCommand{
		householdID: householdID,
		userID:      userID,
	}
}

func (chc *AddUserToHouseholdCommand) Household() uint64 {
	return chc.householdID
}

func (chc *AddUserToHouseholdCommand) User() uint64 {
	return chc.userID
}
