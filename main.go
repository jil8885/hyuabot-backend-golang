package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/bus"
	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/cafeteria"
	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/library"
	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/rest/subway"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

func main() {
	util.ConnectDB()
	app := fiber.New()
	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://prod.backend.hyuabot.app, https://dev.backend.hyuabot.app, https://www.hyuabot.app, " +
			"http://localhost:8100, http://192.168.*.*:8100",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	// RestAPI Routes
	rest := app.Group("/rest")
	bus.SetupRoutes(rest)
	cafeteria.SetupRoutes(rest)
	library.SetupRoutes(rest)
	shuttle.SetupRoutes(rest)
	subway.SetupRoutes(rest)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
