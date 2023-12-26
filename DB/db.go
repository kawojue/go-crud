package db

import (
	"log"

	"github.com/kawojue/go-crud/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := utils.GetEnv("DATABASE_URI", "")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}
