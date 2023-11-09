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
	shuttleGroup.Get("/timetable", v1.GetShuttleTimetableList)
	shuttleGroup.Get("/timetable/:seq", v1.GetShuttleTimetable)
	shuttleGroup.Post("/timetable", v1.CreateShuttleTimetable)
	shuttleGroup.Patch("/timetable/:seq", v1.UpdateShuttleTimetable)
	shuttleGroup.Delete("/timetable/:seq", v1.DeleteShuttleTimetable)
	shuttleGroup.Get("/route", v1.GetShuttleRouteList)
	shuttleGroup.Get("/route/:route", v1.GetShuttleRoute)
	shuttleGroup.Post("/route", v1.CreateShuttleRoute)
	shuttleGroup.Patch("/route/:route", v1.UpdateShuttleRoute)
	shuttleGroup.Delete("/route/:route", v1.DeleteShuttleRoute)
	shuttleGroup.Get("/stop", v1.GetShuttleStopList)
	shuttleGroup.Get("/stop/:stop", v1.GetShuttleStop)
	shuttleGroup.Post("/stop", v1.CreateShuttleStop)
	shuttleGroup.Patch("/stop/:stop", v1.UpdateShuttleStop)
	shuttleGroup.Delete("/stop/:stop", v1.DeleteShuttleStop)
	shuttleGroup.Get("/route/:route/stop", v1.GetShuttleRouteStopList)
	shuttleGroup.Get("/route/:route/stop/:stop", v1.GetShuttleRouteStop)
	shuttleGroup.Post("/route/:route/stop", v1.CreateShuttleRouteStop)
	shuttleGroup.Patch("/route/:route/stop/:stop", v1.UpdateShuttleRouteStop)
	shuttleGroup.Delete("/route/:route/stop/:stop", v1.DeleteShuttleRouteStop)
	shuttleGroup.Get("/period", v1.GetShuttlePeriodList)
	shuttleGroup.Get("/period/:type/:start/:end", v1.GetShuttlePeriod)
	shuttleGroup.Post("/period", v1.CreateShuttlePeriod)
	shuttleGroup.Delete("/period/:type/:start/:end", v1.DeleteShuttlePeriod)
	shuttleGroup.Get("/holiday", v1.GetShuttleHolidayList)
	shuttleGroup.Get("/holiday/:date", v1.GetShuttleHoliday)
	shuttleGroup.Post("/holiday", v1.CreateShuttleHoliday)
	shuttleGroup.Delete("/holiday/:date", v1.DeleteShuttleHoliday)
}
