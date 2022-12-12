package util

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DBInstance struct {
	Database *gorm.DB
}

var instance DBInstance

func ConnectDB() {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_ID"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	instance = DBInstance{Database: db}
}
