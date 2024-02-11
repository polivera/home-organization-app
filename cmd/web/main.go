package main

import (
	"context"
	"fmt"

	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/user/domain/command"
	userservice "github.com/polivera/home-organization-app/internal/user/domain/service"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/repository"
)

func main() {
	var err error
	db := database.NewMySQLConnection(context.Background())
	if err = db.Open(); err != nil {
		panic("can't open database")
	}
	userService := userservice.NewCreateUserService(repository.NewUserRepository(db))
	userDTO, err := userService.Handle(command.NewUserCreateCommand("letest@test.local", "Test.123!", "Testonga"))
	fmt.Println(err, userDTO)
}
