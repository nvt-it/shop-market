package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/shop-market/models"
)

// PGUserRepo explain...
type UserRepo interface {
	Create(ctx *gin.Context, u *models.User) (*models.User, error)
	Fetch(ctx *gin.Context, num int64) ([]*models.User, error)
	// GetByID(ctx *gin.Context, id int64) (*models.User, error)
	// Update(ctx *gin.Context, p *models.User) (*models.User, error)
	// Delete(ctx *gin.Context, id int64) (bool, error)
}

type ShopRepo interface {
	Create(ctx *gin.Context, sh *models.Shop) (*models.Shop, error)
	Fetch(ctx *gin.Context, num int64) ([]*models.Shop, error)
	GetByID(ctx *gin.Context, id int64) (*models.Shop, error)
}

type ProductRepo interface {
	Create(ctx *gin.Context, p *models.Product) (*models.Product, error)
	Fetch(ctx *gin.Context, num int64) ([]*models.Product, error)
}
