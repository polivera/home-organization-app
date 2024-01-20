package pkg

import (
	"context"
	"fmt"
	"github.com/polivera/home-organization-app/pkg/internal/common/infra/database"
	"github.com/polivera/home-organization-app/pkg/internal/user/infra/repository"
)

func Foobar() {
	fmt.Println("I can see here from main")
	ctx := context.Background()
	dbConn := database.NewMySQLConnection(ctx)

	dbConn.Connect()

	repository.GetSomething(dbConn)
}
