package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Shop      Shop      `json:"shop"`
	ShopID    int64     `json:"shop_id"`
}

// Serialize serializes user data to make Restful API
func (p *Product) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"name":       p.Name,
		"price":      p.Price,
		"shop":       p.Shop.Serialize(),
		"created_at": p.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": p.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
