package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"julo_test/service"
	"julo_test/service/model"
	"time"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) service.IWalletRepository {
	return WalletRepository{
		db: db,
	}
}

func (r WalletRepository) Create(ownedBy uuid.UUID) (uuid.UUID, error) {
	wallet := model.Wallet{
		ID:        uuid.New(),
		OwnedBy:   ownedBy,
		Status:    "disabled",
		EnabledAt: &time.Time{},
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	err := r.db.Debug().Table("wallet").Create(&wallet).Error
	if err != nil {
		return uuid.Nil, err
	}

	return wallet.ID, nil
}

func (r WalletRepository) FindWalletByOwnerID(ownerID uuid.UUID) (res model.Wallet, err error) {
	err = r.db.Debug().Table("wallet").Where("owned_by = ?", ownerID.String()).Find(&res).Error
	if err != nil {
		return
	}

	return res, nil
}
