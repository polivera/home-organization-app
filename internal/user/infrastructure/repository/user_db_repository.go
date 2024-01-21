package repository

import (
	"fmt"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/user/domain/repository"
)

type userRepository struct {
	dbConn database.Connection
}

func NewUserRepository(db database.Connection) repository.UserRepository {
	return &userRepository{dbConn: db}
}

func (userRepo *userRepository) GetActiveUserByEmail(email string) {
	fmt.Println(email)
	rows, err := userRepo.dbConn.Query("select * from users")
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}
