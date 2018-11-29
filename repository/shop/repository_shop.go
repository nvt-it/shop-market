package repository

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	models "github.com/shop-market/models"
	repo "github.com/shop-market/repository"
)

// Shop is alias for models.Shop
type Shop = models.Shop

// NewPostgreSQLUserRepo retunrs implement of post repository interface
func NewPGShopRepo(Conn *gorm.DB) repo.ShopRepo {
	return &pgShopRepo{
		Conn: Conn,
	}
}

type pgShopRepo struct {
	Conn *gorm.DB
}

func (db *pgShopRepo) Create(ctx *gin.Context, p *models.Shop) (*models.Shop, error) {

	// create Shop
	shop := Shop{
		Name: p.Name,
	}

	fmt.Println("shop", p)

	isOk := db.Conn.NewRecord(shop)

	if !isOk {
		return nil, errors.New("Record was existed!!!")
	}

	db.Conn.Create(&shop)

	return &shop, nil
}

func (db *pgShopRepo) Fetch(ctx *gin.Context, num int64) ([]*models.Shop, error) {
	shops := make([]*Shop, 0)
	db.Conn.Table("shops").Find(&shops)

	return shops, nil
}

func (db *pgShopRepo) GetByID(ctx *gin.Context, id int64) (*models.Shop, error) {
	shop := Shop{}
	query := fmt.Sprintf("SELECT * FROM shops WHERE id = %s", id)
	db.Conn.Table(query).First(&shop)

	return &shop, nil
}
