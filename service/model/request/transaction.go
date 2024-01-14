package request

import "github.com/google/uuid"

type (
	CreateTransactionRequest struct {
		WalletId        uuid.UUID `json:"wallet_id"`
		Status          string    `json:"status"`
		TransactionType string    `json:"transaction_type"`
		Amount          float64   `json:"amount"`
		RefId           string    `json:"refId"`
	}

	UpdateWalletRequest struct {
		WalletId    uuid.UUID `json:"wallet_id"`
		Amount      float64   `json:"amount"`
		AmountAdded float64   `json:"amount_added"`
		Type        string    `json:"type"`
		RefId       uuid.UUID `json:"refId"`
		DepositBy   uuid.UUID `json:"deposit_by"`
	}
)
