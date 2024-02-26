package command

type GetUserHouseholdsCommand struct {
	userID uint64
}

func NewGetUserHouseholdsCommand(userID uint64) AddUserToHouseholdCommand {
	return AddUserToHouseholdCommand{
		userID: userID,
	}
}

func (chc *GetUserHouseholdsCommand) User() uint64 {
	return chc.userID
}
