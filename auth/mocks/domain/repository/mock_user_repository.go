package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : efishery/auth
 */

import (
	"context"

	dtoLogin "efishery/auth/src/app/dtos/login"
	dto "efishery/auth/src/app/dtos/users"
	models "efishery/auth/src/infra/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (o *MockUserRepo) Create(context context.Context, data *dto.CreateUserReqDTO) (*models.Users, error) {
	args := o.Called(data)

	var (
		err      error
		respData *models.Users
	)
	if n, ok := args.Get(0).(*models.Users); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}

func (o *MockUserRepo) GetUser(context context.Context, data *dtoLogin.LoginReqDTO) (*models.Users, error) {
	args := o.Called(data)

	var (
		err      error
		respData *models.Users
	)
	if n, ok := args.Get(0).(*models.Users); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}
