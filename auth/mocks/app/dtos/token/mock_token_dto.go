package mock_dto

import (
	dto "efishery/auth/src/app/dtos/token"

	"github.com/stretchr/testify/mock"
)

type MockTokenDTO struct {
	mock.Mock
}

func NewMocktokenDTO() *MockTokenDTO {
	return &MockTokenDTO{}
}

var _ dto.TokenDtoInterface = &MockTokenDTO{}

func (m *MockTokenDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
