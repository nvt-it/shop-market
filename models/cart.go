package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type Cart struct {
	ID        int64     `json:"id"`
	ShopID    int64     `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Products  []Product `gorm:"many2many:shop_products;"`
	Shop      Shop
}

// Serialize serializes user data to make Restful API
func (c *Cart) Serialize() common.JSON {
	return common.JSON{
		"id":         c.ID,
		"shop":       c.Shop,
		"created_at": c.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": c.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
