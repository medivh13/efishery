package items_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type TokenDtoInterface interface {
	Validate() error
}

type TokenReqDTO struct {
	Token string `json:"token"`
}

func (dto *TokenReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Token, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type TokenRespDTO struct {
	TokenInfo interface{} `json:"token_info"`
}
