package repository

import (
	"database/sql"
	"errors"
	"github.com/polivera/home-organization-app/internal/household/infrastructure/entity"

	commonValueObject "github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/household/domain/repository"
	"github.com/polivera/home-organization-app/internal/household/domain/valueobject"
)

type householdRepository struct {
	dbConn database.Connection
}

func NewHouseholdRepository(db database.Connection) repository.HouseholdRepository {
	return &householdRepository{dbConn: db}
}

func (h householdRepository) CreateHousehold(householdEntity *entity.Household) error {
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
	name valueobject.HouseholdNameVO,
	owner commonValueObject.IDVO,
) (*entity.Household, error) {
	result := h.dbConn.QueryRow(
		`
		SELECT h.id, h.name, h.owner 
		FROM households h 
		WHERE h.name = ? AND h.owner = ?
		`,
		name.Value(),
		owner.Value(),
	)
	var householdEntity entity.Household
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

func (h householdRepository) GetHouseholdByID(id commonValueObject.IDVO) (*entity.Household, error) {
	result := h.dbConn.QueryRow(
		`
		SELECT h.id, h.name, h.owner 
		FROM households h 
		WHERE h.id = ?
		`,
		id.Value(),
	)
	var householdEntity entity.Household
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
