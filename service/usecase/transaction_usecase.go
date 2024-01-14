package usecase

import (
	"julo_test/service"
	"julo_test/service/model/response"
)

type transactionUsecase struct {
	transactionRepo service.ITransactionRepository
}

func NewTransactionUsecase(transactionRepo service.ITransactionRepository) service.TransactionUsecase {
	return transactionUsecase{
		transactionRepo: transactionRepo,
	}
}

func (t transactionUsecase) FindTransactionByWalletId(walletId string) (res response.SuccessFindTransactionByWalletId, err error) {
	transactionData, err := t.transactionRepo.FindTxByWalletId(walletId)
	if err != nil {
		return res, err
	}

	res.Data.Transactions = transactionData

	return res, nil
}
