package shuttle

import "github.com/gofiber/fiber/v2"

// 셔틀버스 시간표 일괄 조회
func GetShuttleTimeTable(c *fiber.Ctx) error {
	return c.SendString("GetShuttleTimeTable")
}

// 셔틀버스 도착 예정 시간 일괄 조회
func GetShuttleArrivalTime(c *fiber.Ctx) error {
	return c.SendString("GetShuttleArrivalTime")
}
