package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : efishery/auth
 */

import (
	"context"
	dtoLogin "efishery/auth/src/app/dtos/login"
	dto "efishery/auth/src/app/dtos/users"
	models "efishery/auth/src/infra/models"
)

type UserRepository interface {
	Create(context context.Context, data *dto.CreateUserReqDTO) (*models.Users, error)
	GetUser(context context.Context, data *dtoLogin.LoginReqDTO) (*models.Users, error)
}
