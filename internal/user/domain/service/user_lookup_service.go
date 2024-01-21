package service

import (
	"context"
	"fmt"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/repository"
)

func Handle() {
	ctx := context.Background()
	con := database.NewMySQLConnection(ctx)

	err := con.Open()
	if err != nil {
		panic(err)
	}

	repo := repository.NewUserRepository(con)
	data, _ := repo.GetVerifiedUserByEmail("test2@testmail.local")

	fmt.Println(data)
}
