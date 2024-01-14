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
		ID:         uuid.New(),
		OwnedBy:    ownedBy,
		Status:     "disabled",
		EnabledAt:  &time.Time{},
		Balance:    0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
		DisabledAt: nil,
	}

	if err := r.db.Debug().Table("wallet").Create(&wallet).Error; err != nil {
		return uuid.Nil, err
	}

	return wallet.ID, nil
}

func (r WalletRepository) FindWalletByOwnerID(ownerID uuid.UUID) (res model.Wallet, err error) {
	if err = r.db.Debug().Table("wallet").Where("owned_by = ?", ownerID.String()).Find(&res).Error; err != nil {
		return
	}

	return res, nil
}

func (r WalletRepository) FindWalletByWalletID(walletId uuid.UUID) (res model.Wallet, err error) {
	if err = r.db.Debug().Table("wallet").Where("id = ?", walletId.String()).Find(&res).Error; err != nil {
		return
	}

	return res, nil
}

func (r WalletRepository) EnableWallet(walletId uuid.UUID) (res model.Wallet, err error) {
	if err = r.db.Debug().Table("wallet").Updates(map[string]interface{}{
		"status":     "enabled",
		"enabled_at": time.Now(),
	}).Where("id = ?", walletId).Error; err != nil {
		return res, err
	}

	if err = r.db.Debug().Table("wallet").Select("id, owned_by, status, enabled_at, balance").Where("id = ?", walletId).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r WalletRepository) DisableWallet(walletId uuid.UUID) (res model.Wallet, err error) {
	if err = r.db.Debug().Table("wallet").Updates(map[string]interface{}{
		"status":      "disabled",
		"disabled_at": time.Now(),
	}).Where("id = ?", walletId).Error; err != nil {
		return res, err
	}

	if err = r.db.Debug().Table("wallet").Select("id, owned_by, status, enabled_at, balance").Where("id = ?", walletId).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r WalletRepository) UpdateWallet(amount float64) error {
	if err := r.db.Debug().Table("wallet").Updates(map[string]interface{}{
		"balance":    amount,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}
