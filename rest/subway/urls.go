package subway

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	subway := router.Group("/subway")
	subway.Get("/station", GetStationList)
	subway.Get("/station/:station_id", GetStationItem)
	subway.Get("/station/:station_id/arrival", GetStationArrival)
	subway.Get("/station/:station_id/timetable", GetStationTimeTable)
}
