package service

import (
	"github.com/google/uuid"
	"julo_test/service/model"
	"julo_test/service/model/response"
)

type AccountUsecase interface {
	Create() error
	FindByID(id string) (res model.Account, err error)
}

type TransactionUsecase interface {
}

type WalletUsecase interface {
	InitWallet(accountId uuid.UUID) (res response.SuccessInitWallet, err error)
}
