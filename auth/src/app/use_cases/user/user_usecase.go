package user_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth
 */

import (
	"context"
	dtoLogin "efishery/auth/src/app/dtos/login"
	dtoToken "efishery/auth/src/app/dtos/token"
	dto "efishery/auth/src/app/dtos/users"
	"efishery/auth/src/domain/repositories"
	helpers "efishery/auth/src/infra/helpers"
	"errors"
	"log"
)

type UserUsecaseInterface interface {
	Create(ctx context.Context, data *dto.CreateUserReqDTO) (*dto.UserRespDTO, error)
	CreateToken(ctx context.Context, data *dtoLogin.LoginReqDTO) (*dtoLogin.LoginRespDTO, error)
	InfoToken(ctx context.Context, data *dtoToken.TokenReqDTO) (*dtoToken.TokenRespDTO, error)
}

type userUseCase struct {
	UserRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *userUseCase {
	return &userUseCase{
		UserRepo: userRepo,
	}
}

func (uc *userUseCase) Create(ctx context.Context, data *dto.CreateUserReqDTO) (*dto.UserRespDTO, error) {

	dataCreate, err := uc.UserRepo.Create(ctx, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.ToReturnSaveUser(dataCreate), nil
}

func (uc *userUseCase) CreateToken(ctx context.Context, data *dtoLogin.LoginReqDTO) (*dtoLogin.LoginRespDTO, error) {

	dataUser, err := uc.UserRepo.GetUser(ctx, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	token := helpers.GetToken(dataUser.Name, dataUser.Phone, dataUser.Role)

	dataResp := dtoLogin.LoginRespDTO{
		Token: token,
	}

	return &dataResp, nil
}

func (uc *userUseCase) InfoToken(ctx context.Context, data *dtoToken.TokenReqDTO) (*dtoToken.TokenRespDTO, error) {

	dataTkn, status := helpers.ExtractClaims(data.Token)

	if !status {
		return nil, errors.New("invalid token")
	}
	dataResp := dtoToken.TokenRespDTO{
		TokenInfo: dataTkn,
	}
	return &dataResp, nil
}
