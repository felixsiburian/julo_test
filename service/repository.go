package service

import (
	"github.com/google/uuid"
	"julo_test/service/model"
)

type IAccountRepository interface {
	Create(id uuid.UUID) error
	FindByID(id string) (res model.Account, err error)
}

type ITransactionRepository interface {
}

type IWalletRepository interface {
	Create(ownedBy uuid.UUID) (uuid.UUID, error)
	FindWalletByOwnerID(ownerID uuid.UUID) (res model.Wallet, err error)
	FindWalletByWalletID(walletId uuid.UUID) (res model.Wallet, err error)
	EnableWallet(walletId uuid.UUID) (res model.Wallet, err error)
}
