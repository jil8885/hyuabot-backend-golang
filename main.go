package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/hyuabot-developers/hyuabot-backend-golang/api/route"
	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	_ "github.com/hyuabot-developers/hyuabot-backend-golang/docs"
)

func init() {
	database.ConnectDB(os.Getenv("DB_NAME"))
	database.ConnectRedis()
}

func main() {
	app := fiber.New()
	// Logger
	app.Use(logger.New())
	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	// Routes
	route.SetupRouterV1(app)
	// Start server
	log.Fatal(app.Listen(":3000"))
}
