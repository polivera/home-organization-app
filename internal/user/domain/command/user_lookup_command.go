package command

type UserLookupCommand struct {
	email    string
	password string
}

func NewUserLookupCommand(email string, password string) UserLookupCommand {
	return UserLookupCommand{
		email:    email,
		password: password,
	}
}

func (ulc *UserLookupCommand) Email() string {
	return ulc.email
}

func (ulc *UserLookupCommand) Password() string {
	return ulc.password
}
