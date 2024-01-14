package request

type (
	CreateTransactionRequest struct {
		Status          string  `json:"status"`
		TransactionType string  `json:"transaction_type"`
		Amount          float64 `json:"amount"`
	}
)
