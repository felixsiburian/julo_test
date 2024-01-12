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

func (r AccountRepository) Create() error {
	account := model.Account{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Time{},
	}

	err := r.db.Debug().Table("account").Create(&account).Error
	if err != nil {
		return err
	}

	return nil
}
