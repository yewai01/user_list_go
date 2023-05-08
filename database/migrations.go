package database

import (
	"user_list_go/app/models"
	"user_list_go/bootstrap"
)

func Migrate() {
	bootstrap.DB.AutoMigrate(&models.User{})
}
