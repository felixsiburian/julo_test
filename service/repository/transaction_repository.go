package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"julo_test/service"
	"julo_test/service/model"
	"julo_test/service/model/request"
	"time"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) service.ITransactionRepository {
	return TransactionRepository{
		db: db,
	}
}

func (r TransactionRepository) FindTxByWalletId(walletId string) (res []model.Transaction, err error) {
	if err := r.db.Table("transactions").Where("wallet_id = ?", walletId).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r TransactionRepository) CreateTransactions(params request.CreateTransactionRequest) error {
	uuidRef := uuid.MustParse(params.RefId)
	transaction := model.Transaction{
		ID:           uuid.New(),
		Status:       params.Status,
		TransactedAt: time.Now(),
		Type:         params.TransactionType,
		Amount:       params.Amount,
		ReferenceId:  uuidRef,
		WalletId:     params.WalletId,
	}

	if err := r.db.Debug().Table("transactions").Create(&transaction).Error; err != nil {
		fmt.Println("err: ", err)
		return err
	}

	return nil
}
