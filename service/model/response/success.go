package response

import (
	"github.com/google/uuid"
	"julo_test/service/model"
	"time"
)

type (
	SuccessInitWallet struct {
		Data   DataSuccessInitWallet `json:"data"`
		Status string                `json:"status"`
	}

	DataSuccessInitWallet struct {
		Token string `json:"token"`
	}

	SuccessEnableWallet struct {
		Status string                  `json:"status"`
		Data   DataSuccessEnableWallet `json:"data"`
	}

	DataSuccessEnableWallet struct {
		Wallet model.Wallet `json:"wallet"`
	}

	SuccessFindTransactionByWalletId struct {
		Status string                         `json:"status"`
		Data   DataFindTransactionsByWalletId `json:"data"`
	}

	DataFindTransactionsByWalletId struct {
		Transactions []model.Transaction `json:"transactions"`
	}

	SuccessUpdateWallet struct {
		Status string                  `json:"status"`
		Data   DataSuccessUpdateWallet `json:"data"`
	}

	DataSuccessUpdateWallet struct {
		Deposit Deposit `json:"Deposit"`
	}

	Deposit struct {
		ID          uuid.UUID `json:"id"`
		DepositedBy uuid.UUID `json:"deposited_by"`
		Status      string    `json:"status"`
		DepositedAt time.Time `json:"deposited_at"`
		Amount      float64   `json:"amount"`
		ReferenceId uuid.UUID `json:"reference_id"`
	}
)
