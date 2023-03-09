package database

import (
	"go-fiber-tutor/app/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbUrl := "postgres://pg:pass@localhost:5432/crud"
	var err error

	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal(err)
	}

}
