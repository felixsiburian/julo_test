package service

import (
	"github.com/google/uuid"
	"julo_test/service/model/request"
	"julo_test/service/model/response"
)

type AccountUsecase interface {
}

type TransactionUsecase interface {
	FindTransactionByWalletId(walletId string) (res response.SuccessFindTransactionByWalletId, err error)
}

type WalletUsecase interface {
	InitWallet(accountId uuid.UUID) (res response.SuccessInitWallet, err error)
	EnableWallet(token string) (res response.SuccessEnableWallet, err error)
	FindWalletByWalletID(walletId string) (res response.SuccessEnableWallet, err error)
	Deposit(params request.UpdateWalletRequest) (res response.DataSuccessUpdateWallet, err error)
}
