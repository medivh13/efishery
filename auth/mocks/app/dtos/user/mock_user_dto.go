package mock_dto

import (
	dto "efishery/auth/src/app/dtos/users"

	"github.com/stretchr/testify/mock"
)

type MockUserDTO struct {
	mock.Mock
}

func NewMockUserDTO() *MockUserDTO {
	return &MockUserDTO{}
}

var _ dto.UserDtoInterface = &MockUserDTO{}

func (m *MockUserDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
