package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type Order struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	ShopID    int64          `json:"shop_id"`
	Hash      string         `json:"hash"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Histories []OrderHistory `json:"histories"`
	Products  []Product      `gorm:"many2many:shop_products;"`
	Shop      Shop
}

// Serialize serializes user data to make Restful API
func (o *Order) Serialize() common.JSON {
	return common.JSON{
		"id":         o.ID,
		"name":       o.Name,
		"shop_id":    o.ShopID,
		"shop":       o.Shop.Serialize(),
		"created_at": o.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": o.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
