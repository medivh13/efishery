package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth
 */

import (
	"context"
	"errors"

	dtoLogin "efishery/auth/src/app/dtos/login"
	dto "efishery/auth/src/app/dtos/users"
	repositories "efishery/auth/src/domain/repositories"
	helpers "efishery/auth/src/infra/helpers"
	models "efishery/auth/src/infra/models"

	"gorm.io/gorm"
)

type usersRepository struct {
	connection *gorm.DB
}

func NewUsersRepository(db *gorm.DB) repositories.UserRepository {
	return &usersRepository{
		connection: db,
	}
}

func (repo *usersRepository) Create(ctx context.Context, data *dto.CreateUserReqDTO) (*models.Users, error) {
	var userModel models.Users
	q := repo.connection.WithContext(ctx)

	result := q.Raw(`
	SELECT id, name, phone, role, password
	from efishery.users
	WHERE name = ? AND phone = ? AND deleted_at is null`, data.Name, data.Phone).Find(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		password := helpers.RandPass(4)
		if err := q.Raw("INSERT INTO efishery.users (name, phone, role, password) VALUES (?,?,?,?) returning *", data.Name, data.Phone, data.Role, password).Scan(&userModel).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("user already exist")
	}
	return &userModel, nil
}

func (repo *usersRepository) GetUser(ctx context.Context, data *dtoLogin.LoginReqDTO) (*models.Users, error) {
	var userModel models.Users
	q := repo.connection.WithContext(ctx)

	result := q.Raw(`
	SELECT id, name, phone, role, password
	from efishery.users
	WHERE phone = ? AND password = ? AND deleted_at is null`, data.Phone, data.Password).Find(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("invalid user")
	}

	return &userModel, nil
}
