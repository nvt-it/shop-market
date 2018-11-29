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
	user "github.com/shop-market/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repository.UserRepo
}

// NewPostHandler ...
func Authority(db *gorm.DB) *User {
	return &User{
		repo: user.NewPGUserRepo(db),
	}
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func paramsPermited(c *gin.Context) {

}

// Authenticate/Register

func (u *User) Register(c *gin.Context) {
	valid := validation.Validation{}

	passwd, _ := hash(c.Query("Password"))

	// create user
	attributes := models.User{
		Username:     c.Query("Username"),
		PasswordHash: passwd,
	}

	ok, _ := valid.Valid(&attributes)

	if !ok {
		c.JSON(http.StatusBadRequest, common.JSON{"error": "Invalid Data"})

		return
	}

	user, _ := u.repo.Create(c, &attributes)

	serialized := user.Serialize()
	token, _ := security.GenerateToken(serialized)

	c.JSON(http.StatusCreated, common.JSON{
		"message": "Successfully Created",
		"user":    user.Serialize(),
		"token":   token,
	})

	return
}
