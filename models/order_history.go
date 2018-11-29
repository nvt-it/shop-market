package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type OrderHistory struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User
	Order     Order
}

// Serialize serializes user data to make Restful API
func (o *OrderHistory) Serialize() common.JSON {
	return common.JSON{
		"id":         o.ID,
		"order":      o.Order.Serialize(),
		"user":       o.User.Serialize(),
		"created_at": o.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": o.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
