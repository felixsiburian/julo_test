package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"julo_test/service"
	"julo_test/service/model"
	"time"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewServiceCenterRepository(db *gorm.DB) service.IAccountRepository {
	return AccountRepository{
		db: db,
	}
}

func (r AccountRepository) Create(id uuid.UUID) error {
	account := model.Account{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := r.db.Debug().Table("account").Create(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (r AccountRepository) FindByID(id string) (res model.Account, err error) {
	err = r.db.Debug().Table("account").Where("id = ?", id).Find(&res).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return res, err
	}

	return res, nil
}
