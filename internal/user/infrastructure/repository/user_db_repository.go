package repository

import (
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

func (userRepo *userRepository) GetVerifiedUserByEmail(email string) (entity.UserEntity, error) {
	var err error
	result := userRepo.dbConn.QueryRow(
		`
			SELECT u.id, u.email, u.password, u.username, u.session_token, u.status 
			FROM users u 
			WHERE email = ? AND status = ?
		`,
		email,
		entity.StatusVerified,
	)
	if err != nil {
		return entity.UserEntity{}, err
	}

	var userEntity entity.UserEntity
	if err = result.Scan(
		&userEntity.Id,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Username,
		&userEntity.SessionToken,
		&userEntity.Status,
	); err != nil {
		return entity.UserEntity{}, err
	}

	return userEntity, nil
}
