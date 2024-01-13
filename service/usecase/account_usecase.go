package usecase

import (
	"github.com/google/uuid"
	"julo_test/service"
	"julo_test/service/model"
)

type accountUsecase struct {
	accountRepo service.IAccountRepository
}

func NewAccountUsecase(accountRepo service.IAccountRepository) service.AccountUsecase {
	return accountUsecase{
		accountRepo: accountRepo,
	}
}

func (a accountUsecase) Create() error {
	return a.accountRepo.Create(uuid.New())
}

func (a accountUsecase) FindByID(id string) (res model.Account, err error) {
	return a.accountRepo.FindByID(id)
}
