package repository

import (
	"fmt"
	"github.com/polivera/home-organization-app/pkg/internal/common/infra/database"
)

func GetSomething(dbConn database.Connection) {
	rows, err := dbConn.Query("select * from users")
	fmt.Println(rows, err)
}
