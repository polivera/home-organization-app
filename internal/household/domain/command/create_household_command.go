package command

type CreateHouseholdCommand struct {
	name  string
	owner uint64
}

func NewCreateHouseholdCommand(name string, owner uint64) CreateHouseholdCommand {
	return CreateHouseholdCommand{
		name:  name,
		owner: owner,
	}
}

func (chc *CreateHouseholdCommand) Name() string {
	return chc.name
}

func (chc *CreateHouseholdCommand) Owner() uint64 {
	return chc.owner
}
