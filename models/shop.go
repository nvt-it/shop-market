package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type Shop struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Products  []Product `gorm:"many2many:shop_products;"`
}

// Serialize serializes user data to make Restful API
func (s *Shop) Serialize() common.JSON {
	return common.JSON{
		"id":         s.ID,
		"name":       s.Name,
		"products":   s.Products,
		"created_at": s.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": s.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
