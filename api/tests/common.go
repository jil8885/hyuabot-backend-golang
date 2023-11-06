package tests

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/hyuabot-developers/hyuabot-backend-golang/api/route"
	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
)

func setup() *fiber.App {
	app := fiber.New()
	route.SetupRouterV1(app)
	return app
}

func setupDatabase() {
	database.ConnectDB(os.Getenv("DB_NAME"))
	database.ConnectRedis()
}
