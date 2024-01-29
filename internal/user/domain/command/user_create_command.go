package command

type UserCreateCommand struct {
	email    string
	password string
	username string
}

func NewUserCreateCommand(email string, password string, username string) UserCreateCommand {
	return UserCreateCommand{
		email:    email,
		password: password,
		username: username,
	}
}

func (ucc *UserCreateCommand) Email() string {
	return ucc.email
}

func (ucc *UserCreateCommand) Password() string {
	return ucc.password
}

func (ucc *UserCreateCommand) Username() string {
	return ucc.username
}
