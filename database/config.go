package database

import (
	"challenge-week-one/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	connectionString := "host=localhost user=root password=root dbname=challenge port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Error connecting database " + err.Error())
	}

	DB.AutoMigrate(&models.Declaration{})
}
