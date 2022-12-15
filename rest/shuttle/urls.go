package shuttle

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	shuttle := router.Group("/shuttle")
	shuttle.Get("/timetable", GetShuttleTimeTable)
	shuttle.Get("/arrival", GetShuttleArrivalTime)

	route := shuttle.Group("/route")
	route.Get("/", GetShuttleRouteList)
	route.Get("/:route_id", GetShuttleRouteItem)
	route.Get("/:route_id/location", GetShuttleRouteLocation)

	stop := shuttle.Group("/stop")
	stop.Get("/", GetShuttleStopList)
	stop.Get("/:stop_id", GetShuttleStopItem)
	stop.Get("/:stop_id/route/:route_id", GetShuttleStopRoute)
	stop.Get("/:stop_id/route/:route_id/timetable", GetShuttleStopRouteTimeTable)
	stop.Get("/:stop_id/route/:route_id/arrival", GetShuttleStopRouteArrivalTime)
}
