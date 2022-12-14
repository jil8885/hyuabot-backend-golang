package bus

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	// 노선버스 관련 URL 등록
	bus := router.Group("/bus")

	stop := bus.Group("/stop")
	stop.Get("/", GetBusStopList)
	stop.Get("/:stop_id", GetBusStopItem)

	route := bus.Group("/route")
	route.Get("/", GetBusRouteList)
	route.Get("/:route_id", GetBusRouteItem)
	route.Get("/:route_id/timetable/:stop_id", GetBusRouteTimeTable)
}
