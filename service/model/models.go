package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	Account struct {
		ID        uuid.UUID  `json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	Wallet struct {
		ID         uuid.UUID  `gorm:"id" json:"id"`
		OwnedBy    uuid.UUID  `gorm:"owned_by" json:"owned_by"`
		Status     string     `gorm:"status" json:"status"`
		EnabledAt  *time.Time `gorm:"enabled_at" json:"enabled_at"`
		Balance    float64    `gorm:"balance" json:"balance"`
		CreatedAt  time.Time  `gorm:"created_at" json:"created_at"`
		UpdatedAt  time.Time  `gorm:"updated_at" json:"updated_at"`
		DeletedAt  *time.Time `gorm:"deleted_at" json:"deleted_at"`
		DisabledAt *time.Time `gorm:"disabled_at" json:"disabled_at"`
	}

	Transaction struct {
		ID           uuid.UUID `gorm:"id" json:"id"`
		WalletId     uuid.UUID `json:"wallet_id"`
		Status       string    `json:"status"`
		TransactedAt time.Time `json:"transacted_at"`
		Type         string    `json:"type"`
		Amount       float64   `json:"amount"`
		ReferenceId  uuid.UUID `json:"reference_id"`
	}
)
