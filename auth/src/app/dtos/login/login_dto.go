package items_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginDtoInterface interface {
	Validate() error
}

type LoginReqDTO struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (dto *LoginReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Password, validation.Required),
		validation.Field(&dto.Phone, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type LoginRespDTO struct {
	Token string `json:"token"`
}
