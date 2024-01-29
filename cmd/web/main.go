package main

import (
	"context"
	"fmt"
	"github.com/polivera/home-organization-app/internal/user/domain/command"

	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	userservice "github.com/polivera/home-organization-app/internal/user/domain/service"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/repository"
)

func main() {
	var err error
	db := database.NewMySQLConnection(context.Background())
	if err = db.Open(); err != nil {
		panic("can't open database")
	}
	userService := userservice.NewLookupService(repository.NewUserRepository(db))
	userDTO, err := userService.Handle(command.NewUserLookupCommand("test2@testmail.local", "Test.123!"))
	fmt.Println(err, userDTO)
}
