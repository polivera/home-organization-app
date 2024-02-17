package main

import (
	"context"
	"fmt"

	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	"github.com/polivera/home-organization-app/internal/household/domain/service"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/repository"
)

func main() {
	var err error
	db := database.NewMySQLConnection(context.Background())
	if err = db.Open(); err != nil {
		panic("can't open database")
	}

	repo := repository.NewHouseholdRepository(db)
	srv := service.NewCreateHouseholdService(repo)
	dto, err := srv.Handle(command.NewCreateHouseholdCommand("MyHousehold", 9))
	if err != nil {
		fmt.Printf("I got an error: %s", err.Error())
	}

	fmt.Println(dto)
}
