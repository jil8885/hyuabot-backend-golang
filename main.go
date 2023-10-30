package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/hyuabot-developers/hyuabot-backend-golang/api/route"
	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
)

func init() {
	database.ConnectDB()
	database.ConnectRedis()
}

func main() {
	app := fiber.New()
	// Logger
	app.Use(logger.New())
	// Routes
	route.SetupRouterV1(app)
	// Start server
	log.Fatal(app.Listen(":3000"))
}
