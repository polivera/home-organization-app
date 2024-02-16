package household

import "fmt"

type ErrorHouseholdExist struct {
	Name  string
	Owner uint64
}

func (e ErrorHouseholdExist) Error() string {
	return fmt.Sprintf("Household named %s already exist for user %d", e.Name, e.Owner)
}
