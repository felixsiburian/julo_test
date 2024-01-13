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
		ID        uuid.UUID  `gorm:"id" json:"id"`
		OwnedBy   uuid.UUID  `gorm:"owned_by" json:"owned_by"`
		Status    string     `gorm:"status" json:"status"`
		EnabledAt *time.Time `gorm:"enabled_at" json:"enabled_at"`
		Balance   float64    `gorm:"balance" json:"balance"`
		CreatedAt time.Time  `gorm:"created_at" json:"created_at"`
		UpdatedAt time.Time  `gorm:"updated_at" json:"updated_at"`
		DeletedAt *time.Time `gorm:"deleted_at" json:"deleted_at"`
	}
)
