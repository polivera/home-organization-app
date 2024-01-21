package service

import (
	"context"
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
	repo.GetActiveUserByEmail("test2@testmail.local")

}
