package repository

import (
	"database/sql"
	"errors"

	commonValueObject "github.com/polivera/home-organization-app/internal/common/domain/valueobject"
	"github.com/polivera/home-organization-app/internal/common/infrastructure/database"
	"github.com/polivera/home-organization-app/internal/user/domain/repository"
	"github.com/polivera/home-organization-app/internal/user/infrastructure/entity"
)

type userRepository struct {
	dbConn database.Connection
}

func NewUserRepository(db database.Connection) repository.UserRepository {
	return &userRepository{dbConn: db}
}

func (userRepo *userRepository) GetVerifiedUserByEmail(email commonValueObject.EmailVO) (*entity.UserEntity, error) {
	var err error
	result := userRepo.dbConn.QueryRow(
		`
			SELECT u.id, u.email, u.password, u.username, u.session_token, u.status 
			FROM users u 
			WHERE email = ? AND status = ?
		`,
		email.Value(),
		entity.StatusVerified,
	)

	var userEntity entity.UserEntity
	var nullableSessionToken *string
	if err = result.Scan(
		&userEntity.Id,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Username,
		&nullableSessionToken,
		&userEntity.Status,
	); err != nil {
		return nil, err
	}

	if nullableSessionToken != nil {
		userEntity.SessionToken = *nullableSessionToken
	}

	return &userEntity, nil
}

func (userRepo *userRepository) GetUserByID(id commonValueObject.IDVO) (*entity.UserEntity, error) {
	var err error
	result := userRepo.dbConn.QueryRow(
		`
			SELECT u.id, u.email, u.password, u.username, u.session_token, u.status 
			FROM users u 
			WHERE id = ?
		`,
		id.Value(),
	)

	var userEntity entity.UserEntity
	var nullableSessionToken *string
	if err = result.Scan(
		&userEntity.Id,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Username,
		&nullableSessionToken,
		&userEntity.Status,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if nullableSessionToken != nil {
		userEntity.SessionToken = *nullableSessionToken
	}

	return &userEntity, nil
}

func (userRepo *userRepository) GetUserByEmail(email commonValueObject.EmailVO) (*entity.UserEntity, error) {
	var err error
	result := userRepo.dbConn.QueryRow(
		`
			SELECT u.id, u.email, u.password, u.username, u.session_token, u.status 
			FROM users u 
			WHERE email = ?
		`,
		email.Value(),
	)

	var userEntity entity.UserEntity
	var nullableSessionToken *string
	if err = result.Scan(
		&userEntity.Id,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Username,
		&nullableSessionToken,
		&userEntity.Status,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if nullableSessionToken != nil {
		userEntity.SessionToken = *nullableSessionToken
	}

	return &userEntity, nil
}

func (userRepo *userRepository) CreateUser(user *entity.UserEntity) error {
	query := `
		INSERT INTO users (email, password, username, session_token) VALUES (?, ?, ?, '')
	`
	result, err := userRepo.dbConn.Execute(
		query,
		user.Email,
		user.Password,
		user.Username,
	)
	if err != nil {
		return err
	}

	insertedId, err := result.LastInsertId()
	user.Id = uint64(insertedId)

	return err
}
