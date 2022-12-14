package util

import (
	"fmt"
	"github.com/hyuabot-developers/hyuabot-backend-golang/model/library"
	"github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/model/subway"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

type DBInstance struct {
	Database *gorm.DB
}

var DB DBInstance

func ConnectDB() {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_ID"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	shuttle.SetupDatabase(db)
	library.SetupDatabase(db)
	subway.SetupDatabase(db)
	DB = DBInstance{Database: db}
}

type Tabler interface {
	TableName() string
}
