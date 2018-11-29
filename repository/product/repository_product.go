package repository

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	models "github.com/shop-market/models"
	pRepo "github.com/shop-market/repository"
)

// Product is alias for models.Product
type Product = models.Product
type Shop = models.Shop

// NewPostgreSQLUserRepo retunrs implement of post repository interface
func NewPGProductRepo(Conn *gorm.DB) pRepo.ProductRepo {
	return &pgProductRepo{
		Conn: Conn,
	}
}

type pgProductRepo struct {
	Conn *gorm.DB
}

func (db *pgProductRepo) Create(ctx *gin.Context, p *models.Product) (*models.Product, error) {
	shop := Shop{}
	query := fmt.Sprintf("SELECT * FROM shops WHERE id = %s", p.ShopID)
	db.Conn.Table(query).First(&shop)

	// create product
	product := Product{
		Name:   p.Name,
		Shop:   shop,
		ShopID: shop.ID,
	}

	fmt.Println("product", p)

	isOk := db.Conn.NewRecord(product)

	if !isOk {
		return nil, errors.New("Record was existed!!!")
	}

	db.Conn.Create(&product)

	return &product, nil
}

func (db *pgProductRepo) Fetch(ctx *gin.Context, num int64) ([]*models.Product, error) {
	products := make([]*Product, 0)
	db.Conn.Table("products").Find(&products)

	return products, nil
}
