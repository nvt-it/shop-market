package core

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // configures pg driver
	models "github.com/shop-market/models"
)

// DB ...
type DB struct {
	SQL *gorm.DB
	// Mgo *mgo.database
}

// DBConn ...
var dbConn = &DB{}

// Initialize initializes the database
func SetupDatabase() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbName,
	)

	db, err := gorm.Open("postgres", dbConfig)

	db.LogMode(true) // logs PostgreSQL

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	db.AutoMigrate(
		&models.User{},
		&models.Shop{},
		&models.Order{},
		&models.Cart{},
		&models.OrderHistory{},
		&models.Product{},
		&models.Payment{},
		&models.PaymentHistory{},
	)
	// set up foreign keys
	db.Model(&models.Order{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has beed processed")

	return db, err
}
