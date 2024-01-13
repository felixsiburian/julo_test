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
		Token     string     `json:"token"`
	}

	Wallet struct {
		ID        uuid.UUID  `json:"id"`
		OwnedBy   uuid.UUID  `json:"owned_by"`
		Status    string     `json:"status"`
		EnabledAt *time.Time `json:"enabled_at"`
		Balance   float64    `json:"balance"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)
