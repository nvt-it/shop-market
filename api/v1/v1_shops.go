package api_v1

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/shop-market/libs/common"
	"github.com/shop-market/libs/security"
	models "github.com/shop-market/models"
	repository "github.com/shop-market/repository"
	shop "github.com/shop-market/repository/shop"
)

type Shop struct {
	repo repository.ShopRepo
}

func ShopHandler(db *gorm.DB) *Shop {
	return &Shop{
		repo: shop.NewPGShopRepo(db),
	}
}

// Authenticate/Register

func (sh *Shop) Register(c *gin.Context) {
	valid := validation.Validation{}

	// create user
	attributes := models.Shop{
		Name: c.Query("name"),
	}

	ok, _ := valid.Valid(&attributes)

	if !ok {
		c.JSON(http.StatusBadRequest, common.JSON{"error": "Invalid Data"})

		return
	}

	shop, _ := sh.repo.Create(c, &attributes)

	serialized := shop.Serialize()
	token, _ := security.GenerateToken(serialized)

	c.JSON(http.StatusCreated, common.JSON{
		"message": "Successfully Created",
		"shop":    shop.Serialize(),
		"token":   token,
	})

	return
}
