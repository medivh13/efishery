package items_dto

import (
	models "efishery/auth/src/infra/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserDtoInterface interface {
	Validate() error
}

type CreateUserReqDTO struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

func (dto *CreateUserReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Name, validation.Required),
		validation.Field(&dto.Phone, validation.Required),
		validation.Field(&dto.Role, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type UserRespDTO struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func ToReturnSaveUser(d *models.Users) *UserRespDTO {
	return &UserRespDTO{
		ID:       d.ID,
		Name:     d.Name,
		Phone:    d.Phone,
		Role:     d.Role,
		Password: d.Password,
	}
}
