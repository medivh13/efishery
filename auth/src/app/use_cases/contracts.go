package usecases

import (
	UserUc "efishery/auth/src/app/use_cases/user"
)

type AllUseCases struct {
	UserUsecase UserUc.UserUsecaseInterface
}
