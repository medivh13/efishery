package users_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : efishery/auth
 */

import (
	"encoding/json"
	"net/http"

	dtoLogin "efishery/auth/src/app/dtos/login"
	dtoToken "efishery/auth/src/app/dtos/token"
	dto "efishery/auth/src/app/dtos/users"
	usecases "efishery/auth/src/app/use_cases/user"
	common_error "efishery/auth/src/infra/errors"
	"efishery/auth/src/interface/rest/response"
)

type UserHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	CreateToken(w http.ResponseWriter, r *http.Request)
	InfoToken(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	response response.IResponseClient
	usecase  usecases.UserUsecaseInterface
}

func NewUserHandler(r response.IResponseClient, h usecases.UserUsecaseInterface) UserHandlerInterface {
	return &userHandler{
		response: r,
		usecase:  h,
	}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.CreateUserReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.Create(r.Context(), &postDTO)
	if err != nil {
		if err.Error() == "user already exist" {
			h.response.HttpError(w, common_error.NewError(common_error.DATA_ALREADY_EXIST, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_STATUS_PAGE, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Create User",
		data,
		nil,
	)
}

func (h *userHandler) CreateToken(w http.ResponseWriter, r *http.Request) {

	postDTO := dtoLogin.LoginReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.CreateToken(r.Context(), &postDTO)
	if err != nil {
		if err.Error() == "invalid user" {
			h.response.HttpError(w, common_error.NewError(common_error.USER_INVALID, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_STATUS_PAGE, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Login",
		data,
		nil,
	)
}

func (h *userHandler) InfoToken(w http.ResponseWriter, r *http.Request) {

	postDTO := dtoToken.TokenReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.InfoToken(r.Context(), &postDTO)
	if err != nil {
		if err.Error() == "invalid token" {
			h.response.HttpError(w, common_error.NewError(common_error.USER_INVALID, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_STATUS_PAGE, err))
		return
	}

	h.response.JSON(
		w,
		"Valid Token",
		data,
		nil,
	)
}
