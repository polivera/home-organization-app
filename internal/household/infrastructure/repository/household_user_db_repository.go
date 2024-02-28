package repository

import (
	"database/sql"
	"strings"

	"github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
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

func (h householdUserRepository) GetUserHouseholds(userID valueobject.IDVO) ([]entity.Household, error) {
	rows, err := h.dbConn.Query(
		`
		SELECT h.id, h.name, h.owner 
		FROM households h
		INNER JOIN household_users hu ON h.id = hu.household_id
		WHERE h.owner = ? OR hu.user_id = ?;
		`,
		userID.Value(),
		userID.Value(),
	)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			//todo Log
		}
	}(rows)

	if err != nil {
		return nil, err
	}

	var allResults []entity.Household
	for rows.Next() {
		var householdEntity entity.Household

		if err = rows.Scan(
			&householdEntity.ID,
			&householdEntity.Name,
			&householdEntity.Owner,
		); err != nil {
			// todo Log
			continue
		}
		allResults = append(allResults, householdEntity)
	}

	return allResults, nil
}
