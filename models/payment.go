package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type Payment struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id"`
	UserID    int64     `json:"user_id"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Order     Order
	User      User
	Histories []PaymentHistory `json:"histories"`
}

// Serialize serializes user data to make Restful API
func (pay *Payment) Serialize() common.JSON {
	return common.JSON{
		"id":         pay.ID,
		"status":     pay.Status,
		"user_id":    pay.UserID,
		"order_id":   pay.OrderID,
		"order":      pay.Order.Serialize(),
		"user":       pay.User.Serialize(),
		"created_at": pay.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": pay.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
