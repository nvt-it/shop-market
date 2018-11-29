package models

import (
	"time"

	"github.com/shop-market/libs/common"
)

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"color"`
	PasswordHash string    `json:"hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Serialize serializes user data to make Restful API
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":         u.ID,
		"name":       u.Name,
		"username":   u.Username,
		"created_at": u.CreatedAt.Format("15:04 01-02-2006"),
		"updated_at": u.UpdatedAt.Format("15:04 01-02-2006"),
	}
}
