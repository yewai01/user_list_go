package bootstrap

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Dataa *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")" + "/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	Dataa, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect db", err)
	}
	return Dataa
}
