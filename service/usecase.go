package service

import (
	"github.com/google/uuid"
	"julo_test/service/model"
	"julo_test/service/model/response"
)

type AccountUsecase interface {
}

type TransactionUsecase interface {
}

type WalletUsecase interface {
	InitWallet(accountId uuid.UUID) (res response.SuccessInitWallet, err error)
	EnableWallet(token string) (res response.SuccessEnableWallet, err error)
	FindWalletByWalletID(walletId string) (res model.Wallet, err error)
}
