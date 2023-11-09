package route

import (
	"github.com/gofiber/fiber/v2"

	v1 "github.com/hyuabot-developers/hyuabot-backend-golang/api/controller/v1"
	"github.com/hyuabot-developers/hyuabot-backend-golang/utils"
)

func SetupRouterV1(app *fiber.App) {
	api := app.Group("/api/v1")
	authGroup := api.Group("/auth")
	authGroup.Post("/signup", v1.SignUp)
	authGroup.Post("/login", v1.Login)
	authGroup.Post("/refresh", v1.Refresh)
	authGroup.Post("/logout", utils.AuthMiddleware(), v1.Logout)

	shuttleGroup := api.Group("/shuttle", utils.AuthMiddleware())
	shuttleGroup.Get("/timetable/view", v1.GetShuttleTimetableView)
}
