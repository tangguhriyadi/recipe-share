package infrastructure

import (
	"github.com/tangguhriyadi/recipe-share/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "host=localhost user=postgres password=user_service dbname=user_service port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Users{})
	DB = db
}
