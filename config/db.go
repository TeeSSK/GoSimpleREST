package config

import (
	"github.com/TeeSSK/GoSimpleREST/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(
		postgres.Open(
			"host=localhost user=myuser password=mypassword dbname=mydatabase port=5433 sslmode=disable",
		), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
