package repository

import (
	"strings"

	"github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
)

type householdUserRepository struct {
	dbConn database.Connection
}

func NewHouseholdUserRepository(db database.Connection) repository.HouseholdUsersRepository {
	return &householdUserRepository{dbConn: db}
}

func (h householdUserRepository) AddHouseholdUser(householdID valueobject.IDVO, userID valueobject.IDVO) error {
	_, err := h.dbConn.Execute(
		`
		INSERT INTO household_users(household_id, user_id) VALUES (?, ?);
		`,
		householdID.Value(),
		userID.Value(),
	)
	if err != nil {
		switch true {
		case strings.Contains(err.Error(), "users_fk"):
			return database.ErrorForeignRowNotFound{Table: "users"}
		case strings.Contains(err.Error(), "households_fk"):
			return database.ErrorForeignRowNotFound{Table: "households"}
		case strings.Contains(err.Error(), "household_user_uq"):
			return database.ErrorDuplicateKey{Key: "household_user_uq"}
		}
	}
	return err
}
