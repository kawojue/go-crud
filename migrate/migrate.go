package main

import (
	db "github.com/kawojue/go-crud/DB"
	"github.com/kawojue/go-crud/models"
	"github.com/kawojue/go-crud/utils"
)

func init() {
	envPath := "../.env"
	utils.LoadEnv(&envPath)

	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.Post{})
}
