package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"julo_test/service"
	"julo_test/service/model/response"
	"julo_test/service/tools"
)

type walletUsecase struct {
	walletRepo  service.IWalletRepository
	accountRepo service.IAccountRepository
}

func NewWalletUsecase(
	walletRepo service.IWalletRepository,
	accountRepo service.IAccountRepository,
) service.WalletUsecase {
	return walletUsecase{
		walletRepo:  walletRepo,
		accountRepo: accountRepo,
	}
}

func (w walletUsecase) InitWallet(accountId uuid.UUID) (res response.SuccessInitWallet, err error) {
	var walletId uuid.UUID

	accountData, err := w.accountRepo.FindByID(accountId.String())
	if err != nil {
		return
	}

	walletData, err := w.walletRepo.FindWalletByOwnerID(accountId)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}

	if accountData.ID == uuid.Nil {
		err = w.accountRepo.Create(accountId)
		if err != nil {
			return
		}
	}

	if walletData.OwnedBy == uuid.Nil {
		walletId, err = w.walletRepo.Create(accountId)
		if err != nil {
			return
		}

	} else {
		err = errors.New("Already have a wallet")
		return
	}

	token, err := tools.Encrypt(walletId.String())
	if err != nil {
		return
	}

	res.Data.Token = token
	res.Status = "success"

	return res, nil
}

func (w walletUsecase) EnableWallet(walletId string) (res response.SuccessEnableWallet, err error) {
	walletUUID, err := uuid.Parse(walletId)
	if err != nil {
		return
	}

	walletRes, err := w.walletRepo.EnableWallet(walletUUID)
	if err != nil {
		return
	}

	res.Data.Wallet.ID = walletRes.ID
	res.Data.Wallet.Status = walletRes.Status
	res.Data.Wallet.OwnedBy = walletRes.OwnedBy
	res.Data.Wallet.EnabledAt = walletRes.EnabledAt
	res.Data.Wallet.Balance = walletRes.Balance
	res.Status = "success"

	return res, nil
}

func (w walletUsecase) FindWalletByWalletID(walletId string) (res response.SuccessEnableWallet, err error) {
	walletIdParse, err := uuid.Parse(walletId)
	if err != nil {
		return
	}

	walletData, err := w.walletRepo.FindWalletByWalletID(walletIdParse)
	if err != nil {
		return
	}

	res.Data.Wallet = walletData

	return
}
