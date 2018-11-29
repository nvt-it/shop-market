package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type Gen struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Serialize serializes user data to make Restful API
func (g *Gen) Serialize() common.JSON {
	return common.JSON{
		"id":         g.ID,
		"name":       g.Name,
		"color":      g.Color,
		"hash":       g.Hash,
		"created_at": g.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": g.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
