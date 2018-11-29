package repository

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	models "github.com/shop-market/models"
	repo "github.com/shop-market/repository"
)

// User is alias for models.User
type User = models.User

// NewPostgreSQLUserRepo retunrs implement of post repository interface
func NewPGUserRepo(Conn *gorm.DB) repo.UserRepo {
	return &pgUserRepo{
		Conn: Conn,
	}
}

type pgUserRepo struct {
	Conn *gorm.DB
}

func (db *pgUserRepo) Create(ctx *gin.Context, u *models.User) (*models.User, error) {

	// create user
	user := User{
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}

	fmt.Println("user", u)

	isOk := db.Conn.NewRecord(user)

	if !isOk {
		return nil, errors.New("Record was existed!!!")
	}

	db.Conn.Create(&user)

	return &user, nil
}

func (db *pgUserRepo) Fetch(ctx *gin.Context, num int64) ([]*models.User, error) {
	users := make([]*User, 0)
	db.Conn.Table("users").Find(&users)

	return users, nil
}
