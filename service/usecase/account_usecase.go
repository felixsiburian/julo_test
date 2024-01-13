package usecase

import (
	"julo_test/service"
)

type accountUsecase struct {
	accountRepo service.IAccountRepository
}

func NewAccountUsecase(accountRepo service.IAccountRepository) service.AccountUsecase {
	return accountUsecase{
		accountRepo: accountRepo,
	}
}
