package api_v1

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/shop-market/libs/common"
	models "github.com/shop-market/models"
	repository "github.com/shop-market/repository"
	product "github.com/shop-market/repository/product"
)

type Product struct {
	repo repository.ProductRepo
}

// NewProductHandler ...
func ProductHandler(db *gorm.DB) *Product {
	return &Product{
		repo: product.NewPGProductRepo(db),
	}
}

// ProductHandler/Create

func (p *Product) Create(c *gin.Context) {
	valid := validation.Validation{}

	id := c.Query("shop_id")
	shop_id, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.JSON{"error": "Invalid Shop"})

		return
	}

	// create product
	attributes := models.Product{
		Name:   c.Query("name"),
		ShopID: shop_id,
	}

	ok, _ := valid.Valid(&attributes)

	if !ok {
		c.JSON(http.StatusBadRequest, common.JSON{"error": "Invalid Data"})

		return
	}

	product, _ := p.repo.Create(c, &attributes)
	serialized := product.Serialize()

	c.JSON(http.StatusCreated, common.JSON{
		"message": "Successfully Created",
		"product": serialized,
	})

	return
}
