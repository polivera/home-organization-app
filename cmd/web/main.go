package main

import (
	"context"
	"fmt"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/household/domain/command"
	"github.com/polivera/home-organization-app/internal/household/domain/service"
	householdRepository "github.com/polivera/home-organization-app/internal/household/infrastructure/repository"
	userRepository "github.com/polivera/home-organization-app/internal/user/infrastructure/repository"
)

func main() {
	var err error
	db := database.NewMySQLConnection(context.Background())
	if err = db.Open(); err != nil {
		panic("can't open database")
	}

	householdRepo := householdRepository.NewHouseholdRepository(db)
	householdUserRepo := householdRepository.NewHouseholdUserRepository(db)
	userRepo := userRepository.NewUserRepository(db)

	huService := service.NewAddHouseholdUserService(householdRepo, householdUserRepo, userRepo)
	res, err := huService.Handle(command.NewAddUserToHouseholdCommand(2, 8))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

// 1, 8
