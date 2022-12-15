package test

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/bus"
	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/cafeteria"
	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/library"
	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/subway"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

func InitApp() *fiber.App {
	// Setup Fiber App
	util.ConnectDB()
	app := fiber.New()
	rest := app.Group("/rest")
	bus.SetupRoutes(rest)
	cafeteria.SetupRoutes(rest)
	library.SetupRoutes(rest)
	shuttle.SetupRoutes(rest)
	subway.SetupRoutes(rest)

	return app
}