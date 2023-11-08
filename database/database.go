package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v7"
	_ "github.com/lib/pq" //nolint:goimports
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB
var Client *redis.Client

func ConnectDB(databaseName string) {
	var connectionURL = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		databaseName,
	)

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	DB = db
}

func ConnectRedis() {
	connectionURL := os.Getenv("REDIS_URL")
	if connectionURL == "" {
		connectionURL = "localhost:6379"
	}
	Client = redis.NewClient(&redis.Options{
		Addr: connectionURL,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
