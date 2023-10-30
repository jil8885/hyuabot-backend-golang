package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	"log"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()
	// Logger
	app.Use(logger.New())
	// Routes
	// Start server
	log.Fatal(app.Listen(":3000"))
}
