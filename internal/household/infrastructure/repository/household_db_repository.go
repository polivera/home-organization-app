package repository

import (
	"database/sql"
	"errors"

	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	commonValueObject "github.com/polivera/home-organization-app/internal/common/valueobject"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
	"github.com/polivera/home-organization-app/internal/household/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"
)

type householdRepository struct {
	dbConn database.Connection
}

func NewHouseholdRepository(db database.Connection) repository.HouseholdRepository {
	return &householdRepository{dbConn: db}
}

func (h householdRepository) CreateHousehold(householdEntity *entity.HouseholdEntity) error {
	result, err := h.dbConn.Execute(
		`
		INSERT INTO households (name, owner) VALUES (?, ?);
		`,
		householdEntity.Name,
		householdEntity.Owner,
	)
	if err != nil {
		return err
	}
	insertedId, err := result.LastInsertId()
	householdEntity.Id = uint64(insertedId)
	return err
}

func (h householdRepository) GetUserHouseholdByName(
	name valueobject.HouseholdName,
	owner commonValueObject.ID,
) (*entity.HouseholdEntity, error) {
	result := h.dbConn.QueryRow(
		`
		SELECT h.id, h.name, h.owner 
		FROM households h 
		WHERE h.name = ? AND h.owner = ?
		`,
		name.Value(),
		owner.Value(),
	)
	var householdEntity entity.HouseholdEntity
	if err := result.Scan(
		&householdEntity.Id,
		&householdEntity.Name,
		&householdEntity.Owner,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &householdEntity, nil
}
