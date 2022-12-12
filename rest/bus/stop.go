package bus

import "github.com/gofiber/fiber/v2"

// 버스 정류장 목록 조회
func GetBusStopList(c *fiber.Ctx) error {
	return c.SendString("GetBusStopList")
}

// 버스 정류장 항목 조회
func GetBusStopItem(c *fiber.Ctx) error {
	return c.SendString("GetBusStopItem")
}
