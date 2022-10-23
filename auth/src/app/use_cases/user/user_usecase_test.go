package user_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : efishery/auth
 */

import (
	"context"
	"errors"
	"testing"
	"time"

	mockDTOLogin "efishery/auth/mocks/app/dtos/login"
	mockDTOToken "efishery/auth/mocks/app/dtos/token"
	mockDTOuser "efishery/auth/mocks/app/dtos/user"
	mockUserRepo "efishery/auth/mocks/domain/repository"

	dtoLogin "efishery/auth/src/app/dtos/login"
	dtoToken "efishery/auth/src/app/dtos/token"
	dtoUser "efishery/auth/src/app/dtos/users"
	models "efishery/auth/src/infra/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseUserTest struct {
	suite.Suite
	userRepo     *mockUserRepo.MockUserRepo
	usecase      UserUsecaseInterface
	userModels   *models.Users
	dtoTest      *dtoUser.CreateUserReqDTO
	dtoTokenFail *dtoToken.TokenReqDTO
	mockDTO      *mockDTOuser.MockUserDTO
	mockDTOLogin *mockDTOLogin.MockLoginDTO
	mockDTOToken *mockDTOToken.MockTokenDTO
	dtoLoginTest *dtoLogin.LoginReqDTO
	dtoTokenTest *dtoToken.TokenReqDTO
}

func (suite *UsecaseUserTest) SetupTest() {
	suite.userRepo = new(mockUserRepo.MockUserRepo)
	suite.mockDTO = new(mockDTOuser.MockUserDTO)
	suite.usecase = NewUserUseCase(suite.userRepo)

	suite.userModels = &models.Users{
		ID:        1,
		Name:      "Test",
		Phone:     "1234567",
		Role:      "admin",
		Password:  "Password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.dtoTest = &dtoUser.CreateUserReqDTO{
		Name:  "Test",
		Phone: "1234567",
		Role:  "admin",
	}

	suite.dtoLoginTest = &dtoLogin.LoginReqDTO{
		Phone:    "1234567",
		Password: "xxddxx",
	}

	suite.dtoTokenTest = &dtoToken.TokenReqDTO{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoidGVzdCIsIlBob25lIjoiMTIzNDU2NyIsIlJvbGUiOiJhZG1pbiIsIlRpbWUiOiIyMDIyLTEwLTIzVDE2OjA0OjE1LjU2NjIyNSswNzowMCJ9.JlqXwvSJDmVHdo-W1tsq8GNhItNbA-V4oqJpy2hcaeI",
	}

	suite.dtoTokenFail = &dtoToken.TokenReqDTO{
		Token: "xxdxx",
	}

}

func (uc *UsecaseUserTest) TestCreateSuccess() {
	uc.userRepo.Mock.On("Create", uc.dtoTest).Return(uc.userModels, nil)
	_, err := uc.usecase.Create(context.Background(), uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseUserTest) TestCreateFail() {
	uc.userRepo.Mock.On("Create", uc.dtoTest).Return(nil, errors.New(mock.Anything))
	_, err := uc.usecase.Create(context.Background(), uc.dtoTest)
	uc.Error(errors.New(mock.Anything), err)
}

func (uc *UsecaseUserTest) TestLoginSuccess() {
	uc.userRepo.Mock.On("GetUser", uc.dtoLoginTest).Return(uc.userModels, nil)
	_, err := uc.usecase.CreateToken(context.Background(), uc.dtoLoginTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseUserTest) TestLoginFail() {
	uc.userRepo.Mock.On("GetUser", uc.dtoLoginTest).Return(nil, errors.New(mock.Anything))
	_, err := uc.usecase.CreateToken(context.Background(), uc.dtoLoginTest)
	uc.Error(errors.New(mock.Anything), err)
}

func (uc *UsecaseUserTest) TestTokenSuccess() {
	_, err := uc.usecase.InfoToken(context.Background(), uc.dtoTokenTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseUserTest) TestTokenFail() {
	_, err := uc.usecase.InfoToken(context.Background(), uc.dtoTokenFail)
	uc.Error(errors.New("invalid token"), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseUserTest))
}
