package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type PaymentHistory struct {
	ID        int64     `json:"id"`
	PaymentID int64     `json:"payment_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User
	Payment   Payment
}

// Serialize serializes user data to make Restful API
func (ph *PaymentHistory) Serialize() common.JSON {
	return common.JSON{
		"id":         ph.ID,
		"payment":    ph.Payment.Serialize(),
		"user":       ph.User.Serialize(),
		"created_at": ph.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": ph.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
