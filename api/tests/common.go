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

func tearDownDatabase() {
	// Remove all bus data
	_, _ = database.DB.Exec("DELETE FROM bus_realtime")
	_, _ = database.DB.Exec("DELETE FROM bus_timetable")
	_, _ = database.DB.Exec("DELETE FROM bus_route_stop")
	_, _ = database.DB.Exec("DELETE FROM bus_route")
	_, _ = database.DB.Exec("DELETE FROM bus_stop")
	// Remove all commute shuttle data
	_, _ = database.DB.Exec("DELETE FROM commute_shuttle_timetable")
	_, _ = database.DB.Exec("DELETE FROM commute_shuttle_route")
	_, _ = database.DB.Exec("DELETE FROM commute_shuttle_stop")
	// Remove all shuttle data
	_, _ = database.DB.Exec("DELETE FROM shuttle_timetable")
	_, _ = database.DB.Exec("DELETE FROM shuttle_route_stop")
	_, _ = database.DB.Exec("DELETE FROM shuttle_route")
	_, _ = database.DB.Exec("DELETE FROM shuttle_stop")
	_, _ = database.DB.Exec("DELETE FROM shuttle_holiday")
	_, _ = database.DB.Exec("DELETE FROM shuttle_period")
	_, _ = database.DB.Exec("DELETE FROM shuttle_period_type")
	// Remove all subway data
	_, _ = database.DB.Exec("DELETE FROM subway_realtime")
	_, _ = database.DB.Exec("DELETE FROM subway_timetable")
	_, _ = database.DB.Exec("DELETE FROM subway_route_station")
	_, _ = database.DB.Exec("DELETE FROM subway_route")
	_, _ = database.DB.Exec("DELETE FROM subway_station")
	// Remove all meal data
	_, _ = database.DB.Exec("DELETE FROM menu")
	_, _ = database.DB.Exec("DELETE FROM restaurant")
	// Remove all reading room data
	_, _ = database.DB.Exec("DELETE FROM reading_room")
	// Remove all campus data
	_, _ = database.DB.Exec("DELETE FROM campus")
	// Remove all notice data
	_, _ = database.DB.Exec("DELETE FROM notices")
	_, _ = database.DB.Exec("DELETE FROM notice_category")
	// Remove all users
	_, _ = database.DB.Exec("DELETE FROM admin_user")
}
